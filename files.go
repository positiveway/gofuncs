package gofuncs

import (
	"os"
)

func ReadFile(file Str) Str {
	data, err := os.ReadFile(file)
	CheckErr(err)
	return Str(data)
}

func ReadLines(file Str) []Str {
	content := ReadFile(file)
	return Split(content, "\n")
}
