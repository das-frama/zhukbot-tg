package txtdb

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
)

type db struct {
	files map[string]*os.File
}

func New(dir string) (db, error) {
	d := db{
		files: map[string]*os.File{
			"chats.txt": nil,
			"users.txt": nil,
			"zhuks.txt": nil,
		},
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

func (db *db) Close() {
	for _, file := range db.files {
		file.Close()
	}
}

func (db *db) Insert(name string, s Stringer) {
	file := db.files[name]
	file.WriteString(s.ToString() + "\n")
	file.Seek(0, 0)
}

func findStructByFile(name string) interface{} {
	if name == "chats.txt" {
		return Chat{}
	} else if name == "users.txt" {
		return User{}
	} else if name == "zhuks.txt" {
		return Zhuk{}
	}
	return nil
}

// func fileNames(mm map[string]*os.File) []string {
// 	keys := make([]string, 0, len(mm))
// 	for m := range mm {
// 		keys = append(keys, m)
// 	}
// 	return keys
// }

func createOrOpenFile(name string) (*os.File, error) {
	if fileExists(name) {
		return os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0755)
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

// func fetchValues(schema interface{}) []string {
// 	var values []string

// 	v := reflect.ValueOf(schema)
// 	for i := 0; i < v.NumField(); i++ {
// 		if v.Field(i).Type().PkgPath() == "" {
// 			value := v.Field(i).Elem().
// 			values = append(values, value.(string))
// 		}
// 	}

// 	return values
// }

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
