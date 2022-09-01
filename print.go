package gofuncs

import "fmt"

var PrintDebugInfo bool

func Format(message string, variables ...any) string {
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

func Print(message string, variables ...any) {
	fmt.Print(Format(message, variables...))
}

func PrintDebug(message string, variables ...any) {
	if PrintDebugInfo {
		Print(message, variables...)
	}
}
