package txtdb

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
)

type DB struct {
	files map[string]*os.File
}

func New(dir string) (DB, error) {
	d := DB{
		files: map[string]*os.File{
			"chats.txt": nil,
			"users.txt": nil,
			"zhuks.txt": nil,
		},
	}

	// Check if directory exists.
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	for name := range d.files {
		var err error
		d.files[name], err = createOrOpenFile(path.Join(dir, name))
		if err != nil {
			return d, err
		}
		// Populate header if empty.
		if bytes, _ := ioutil.ReadAll(d.files[name]); len(bytes) == 0 {
			s := findStructByFile(name)
			d.files[name].WriteString(strings.Join(fetchTypes(s), "\t") + "\n")
		}
	}

	return d, nil
}

func (db *DB) Close() {
	for _, file := range db.files {
		file.Close()
	}
}

func (db *DB) Insert(name string, t Tabler) error {
	// Check if record already exists.
	keyField, keyValue := t.Key()
	if _, err := db.Fetch(name, keyField, keyValue); err != ErrNotFound {
		return ErrAlreadyExists
	}

	file := db.files[name]
	file.Seek(0, 2)

	file.WriteString(t.ToString() + "\n")
	err := file.Sync()

	return err
}

func (db *DB) Update(name string, t Tabler) {
	file := db.files[name]
	file.Seek(0, 0)

	found := false
	writer := bufio.NewWriter(file)
	scanner := bufio.NewScanner(file)
	// Scan header.
	scanner.Scan()
	writer.Write(scanner.Bytes())
	writer.Write([]byte("\n"))

	_, key := t.Key()
	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), "\t")
		if ss[0] == key {
			writer.WriteString(t.ToString() + "\n")
			found = true
		} else {
			writer.Write(scanner.Bytes())
			writer.Write([]byte("\n"))
		}
	}

	if found {
		file.Truncate(0)
		file.Seek(0, 0)
		writer.Flush()
	}
}

func (db *DB) Delete(name string, t Tabler) error {
	file := db.files[name]
	file.Seek(0, 0)

	found := false
	writer := bufio.NewWriter(file)
	scanner := bufio.NewScanner(file)

	// Scan header.
	scanner.Scan()
	writer.Write(scanner.Bytes())
	writer.Write([]byte("\n"))

	_, key := t.Key()
	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), "\t")
		if ss[0] == key {
			found = true
		} else {
			writer.Write(scanner.Bytes())
			writer.Write([]byte("\n"))
		}
	}

	if found {
		file.Truncate(0)
		file.Seek(0, 0)
		writer.Flush()
	}

	return nil
}

func (db *DB) Fetch(name string, key string, value string) (Tabler, error) {
	s := findStructByFile(name)
	file := db.files[name]
	file.Seek(0, 0)

	// Read headers.
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return s, err
	}
	headers := strings.Split(scanner.Text(), "\t")
	if len(headers) == 1 {
		return s, fmt.Errorf("empty header in %s file", file.Name())
	}

	// Find header index.
	idx := search(key, headers)
	if idx == -1 {
		return s, fmt.Errorf("there is no any header in %s file", file.Name())
	}

	// Find the right record.
	record := make([]string, 0, len(headers))
	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), "\t")
		if ss[idx] == value {
			record = ss
			break
		}
	}
	if len(record) == 0 {
		return s, ErrNotFound
	}

	// Explode line to the struct.
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(&s).Elem()
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)

		idx = search(ft.Tag.Get("txtdb"), headers)
		if idx != -1 {
			tmp := reflect.New(v.Elem().Type()).Elem()
			tmp.Set(v.Elem())

			switch ft.Type.Kind() {
			case reflect.String:
				tmp.Field(i).SetString(record[idx])
			case reflect.Int:
				intVal, _ := strconv.ParseInt(record[idx], 10, 64)
				tmp.Field(i).SetInt(intVal)
			case reflect.Bool:
				boolVal, _ := strconv.ParseBool(record[idx])
				tmp.Field(i).SetBool(boolVal)
			}
			v.Set(tmp)
		}

	}

	return s, nil
}

func (db *DB) HasRecord(name string, key string, value string) bool {
	_, err := db.Fetch(name, key, value)
	return err == nil
}

func search(needle string, haystack []string) int {
	for i, h := range haystack {
		if h == needle {
			return i
		}
	}

	return -1
}

func findStructByFile(name string) Tabler {
	if name == "chats.txt" {
		return Chat{}
	} else if name == "users.txt" {
		return User{}
	} else if name == "zhuks.txt" {
		return Zhuk{}
	}
	return nil
}

func createOrOpenFile(name string) (*os.File, error) {
	if fileExists(name) {
		return os.OpenFile(name, os.O_RDWR, 0755)
	}

	return os.Create(name)
}

func fetchTypes(schema interface{}) []string {
	var headers []string

	t := reflect.TypeOf(schema)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).PkgPath == "" {
			headers = append(headers, t.Field(i).Tag.Get("txtdb"))
		}
	}

	return headers
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
