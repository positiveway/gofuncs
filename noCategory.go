package gofuncs

import "github.com/jinzhu/copier"

func Swap[T any](value1, value2 *T) {
	*value1, *value2 = *value2, *value1
}

func Copy[T any](value T) T {
	//Takes long time
	var copiedValue T

	err := copier.Copy(&copiedValue, value)
	if err != nil {
		Panic("Copying failed")
	}

	return copiedValue
}
