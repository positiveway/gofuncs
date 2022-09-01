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

func JoinPathCheckIfExists(elem ...string) string {
	filePath := filepath.Join(elem...)
	CheckFileOrDirExists(filePath)
	return filePath
}

func ReadFile[T EncodingT](elem ...string) T {
	data, err := os.ReadFile(JoinPathCheckIfExists(elem...))
	CheckErr(err)
	return T(data)
}

func ReadFileStr(elem ...string) string {
	return ReadFile[string](elem...)
}

func ReadJson(res interface{}, elem ...string) {
	filePath := JoinPathCheckIfExists(elem...)

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

func ReadLines(elem ...string) []string {
	content := ReadFileStr(elem...)
	return Split(content, "\n")
}

func ReadLayoutFile(skipLines int, elem ...string) [][]string {
	lines := ReadLines(elem...)
	lines = lines[skipLines:]

	var linesParts [][]string
	for _, line := range lines {
		line = Strip(line)
		if IsEmptyStripStr(line) || StartsWithAnyOf(line, ";", "//") {
			continue
		}
		parts := SplitByAnyOf(line, "&|>:,=")
		for ind, part := range parts {
			parts[ind] = Strip(part)
		}
		linesParts = append(linesParts, parts)
	}
	return linesParts
}
