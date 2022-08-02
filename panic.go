package gofuncs

import (
	"log"
)

func GetPanicMsg(message []Str, defaultMsg Str) Str {
	switch len(message) {
	case 0:
		return defaultMsg
	case 1:
		return message[0]
	default:
		Panic("Only one message can be specified")
	}
	panic("")
}

func Panic(message Str, variables ...any) {
	log.Fatal(Format(message, variables...))
}

func PanicErr(err error) {
	Panic("%v", err)
}

func PanicMisspelled(str any) {
	Panic("Probably misspelled: %v", str)
}

func CheckErr(err error) {
	if err != nil {
		PanicErr(err)
	}
}
