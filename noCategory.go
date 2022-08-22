package gofuncs

import "github.com/jinzhu/copier"

func Swap[T any](value1, value2 *T) {
	*value1, *value2 = *value2, *value1
}

func Copy[T any](toValue, fromValue *T) {
	err := copier.Copy(toValue, fromValue)
	if err != nil {
		Panic("Copying failed")
	}
}
