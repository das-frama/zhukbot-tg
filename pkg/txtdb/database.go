package txtdb

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
)

type db struct {
	chatsFile *os.File
	usersFile *os.File
}

func New(dir string) (db, error) {
	d := db{}

	names, err := createFilesIfNotExists(dir, []string{"chats.txt", "users.txt"})
	if err != nil {
		return d, nil
	}

	if d.chatsFile, err = os.OpenFile(names[0], os.O_RDWR, 0755); err != nil {
		return d, err
	}
	if d.usersFile, err = os.OpenFile(names[1], os.O_RDWR, 0755); err != nil {
		return d, err
	}

	if bytes, _ := ioutil.ReadAll(d.chatsFile); len(bytes) == 0 {
		chatTypes := fetchTypes(Chat{})
		d.chatsFile.WriteString(strings.Join(chatTypes, "\t"))
	}
	if bytes, _ := ioutil.ReadAll(d.usersFile); len(bytes) == 0 {
		userTypes := fetchTypes(User{})
		d.usersFile.WriteString(strings.Join(userTypes, "\t"))
	}

	return d, nil
}

func (db *db) Close() {
	db.usersFile.Close()
	db.chatsFile.Close()
}

func createFilesIfNotExists(dir string, names []string) ([]string, error) {
	// list of all files.
	paths := make([]string, len(names))
	for i, name := range names {
		paths[i] = path.Join(dir, name)
		if fileExists(paths[i]) {
			continue
		}

		file, err := os.Create(paths[i])
		if err != nil {
			return paths, err
		}
		file.Close()
	}

	return paths, nil
}

// func hasHeader(file *os.File) bool {
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := strings.TrimSpace(scanner.Text())
// 		if strings.HasPrefix(line, "#") {
// 			continue
// 		}

// 	}
// }

func fetchTypes(schema interface{}) []string {
	var headers []string

	t := reflect.TypeOf(schema)
	for i := 0; i < t.NumField(); i++ {
		headers = append(headers, t.Field(i).Tag.Get("txtdb"))
	}

	return headers
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
