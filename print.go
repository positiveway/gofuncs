package gofuncs

import "fmt"

var PrintDebugInfo bool

func Format(message Str, variables ...any) Str {
	if IsNotInit(message) {
		for i := 0; i < len(variables); i++ {
			message += "%v "
		}
		message = RemoveLastLetters(message, 1)
	}
	if !(StartsWith(message, "\n") || EndsWith(message, "\n")) {
		message += "\n"
	}
	return fmt.Sprintf(message, variables...)
}

func Print(message Str, variables ...any) {
	fmt.Print(Format(message, variables...))
}

func PrintDebug(message Str, variables ...any) {
	if PrintDebugInfo {
		Print(message, variables...)
	}
}
