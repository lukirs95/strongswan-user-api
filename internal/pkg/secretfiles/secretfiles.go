package secretfiles

import (
	"fmt"
	"os"
)

func GetFile(name string) ([]byte, error) {
	file, err := os.ReadFile(fmt.Sprintf("static/secret/%s", name))
	if err != nil {
		return make([]byte, 0), err
	}
	return file, nil
}

func GetDir() (string, error) {
	dir, err := os.ReadDir("static/secret")
	if err != nil {
		return "", err
	}
	fileListJson := "["
	for _, file := range dir {
		fileInfo, _ := file.Info()
		fileListJson = fmt.Sprintf("%s%s, ", fileListJson, fileInfo.Name())
	}
	return fileListJson[:len(fileListJson)-2] + "]", nil
}
