package gofuncs

import "github.com/jinzhu/copier"

func Copy[T any](value T) T {
	//Takes long time
	var copiedValue T

	err := copier.Copy(&copiedValue, value)
	if err != nil {
		Panic("Copying failed")
	}

	return copiedValue
}
