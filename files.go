package gofuncs

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

func GetCurFileDir() string {
	execPath, err := os.Executable()
	CheckErr(err)

	execPath, err = filepath.EvalSymlinks(execPath)
	CheckErr(err)

	dirPath := filepath.Dir(execPath)
	return dirPath
}

type Bytes = []byte

type EncodingT interface {
	Bytes | string
}

func CheckFileOrDirExists(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		Panic("File or folder with such name doesn't exist: %s", filePath)
	}
}

func ReadFile[T EncodingT](filePath string) T {
	data, err := os.ReadFile(filePath)
	CheckErr(err)
	return T(data)
}

func ReadJson(filePath string, res interface{}) {
	if !EndsWith(filePath, ".json") {
		Panic("Invalid file extension")
	}
	rawContent := ReadFile[Bytes](filePath)
	dec := json.NewDecoder(bytes.NewReader(rawContent))
	for {
		if err := dec.Decode(res); err == io.EOF {
			break
		} else {
			CheckErr(err)
		}
	}
}

func ReadLines(file string) []string {
	content := ReadFile[string](file)
	return Split(content, "\n")
}
