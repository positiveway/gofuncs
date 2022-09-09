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

func JoinPathCheckIfExists(pathFragments ...string) string {
	filePath := filepath.Join(pathFragments...)
	CheckFileOrDirExists(filePath)
	return filePath
}

func ReadFile[T EncodingT](pathFragments ...string) T {
	data, err := os.ReadFile(JoinPathCheckIfExists(pathFragments...))
	CheckErr(err)
	return T(data)
}

func ReadFileStr(pathFragments ...string) string {
	return ReadFile[string](pathFragments...)
}

func ReadJson(res interface{}, pathFragments []string) {
	filePath := JoinPathCheckIfExists(pathFragments...)

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

func ReadLines(pathFragments ...string) []string {
	content := ReadFileStr(pathFragments...)
	return Split(content, "\n")
}

func ReadLayoutFile(skipLines int, pathFragments []string) [][]string {
	lines := ReadLines(pathFragments...)
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
