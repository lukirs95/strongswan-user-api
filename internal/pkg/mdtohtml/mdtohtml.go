package mdtohtml

import (
	"fmt"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func GetFile(name string) ([]byte, error) {
	input, err := os.ReadFile(fmt.Sprintf("static/md/%s", name))
	if err != nil {
		input, err = os.ReadFile("static/md/404.md")
		if err != nil {
			panic(err)
		}
	}

	unsafe := blackfriday.Run(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return html, nil
}

func GetDir() (string, error) {
	dir, err := os.ReadDir("static/md")
	if err != nil {
		panic(err)
	}
	fileListJson := "["
	for _, file := range dir {
		fileInfo, _ := file.Info()
		fileListJson = fmt.Sprintf("%s%s, ", fileListJson, fileInfo.Name()[:len(fileInfo.Name())-3])
	}
	return fileListJson[:len(fileListJson)-2] + "]", nil
}
