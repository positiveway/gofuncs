package gofuncs

import (
	"strings"
)

func RemoveLastLetters(sourceStr Str, numLettersToRemove int) Str {
	return sourceStr[:len(sourceStr)-numLettersToRemove]
}

func RemoveFirstLetters(sourceStr Str, numLettersToRemove int) Str {
	return sourceStr[numLettersToRemove:]
}

func SplitByAnyOf(sourceStr Str, separators Str) []Str {
	if separators == "" {
		Panic("Empty separator")
	}
	var res []Str
	prevSplitInd := 0
	for ind, symbol := range sourceStr {
		if strings.ContainsRune(separators, symbol) {
			res = append(res, sourceStr[prevSplitInd:ind])
			prevSplitInd = ind + 1
		}
	}
	prevSplitInd = Min(len(sourceStr), prevSplitInd)
	res = append(res, sourceStr[prevSplitInd:])
	return res
}

func StartsWith(s Str, prefix Str) bool {
	return strings.HasPrefix(s, prefix)
}

func EndsWith(s Str, suffix Str) bool {
	return strings.HasSuffix(s, suffix)
}

func StartsWithAnyOf(s Str, prefixes ...Str) bool {
	for _, prefix := range prefixes {
		if StartsWith(s, prefix) {
			return true
		}
	}
	return false
}

func EndsWithAnyOf(s Str, suffixes ...Str) bool {
	for _, suffix := range suffixes {
		if EndsWith(s, suffix) {
			return true
		}
	}
	return false
}

func TrimAnyPrefix(s Str, prefixes ...Str) Str {
	for _, prefix := range prefixes {
		if StartsWith(s, prefix) {
			return strings.TrimPrefix(s, prefix)
		}
	}
	return s
}

func Strip(s Str) Str {
	return strings.TrimSpace(s)
}

func FilterEmptyStrings(slice []Str) []Str {
	var filtered []Str
	for _, s := range slice {
		if !IsEmptyStripStr(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func Split(s Str, sep ...Str) []Str {
	PanicIfEmptyStripStr(s)

	separator := func(separator []Str) Str {
		switch len(separator) {
		case 0:
			return ""
		case 1:
			return sep[0]
		default:
			Panic("Invalid separator for split")
		}
		panic("")
	}(sep)

	res := strings.Split(s, separator)
	res = FilterEmptyStrings(res)
	return res
}

func Words(s Str) []Str {
	return Split(s)
}

func FirstWord(s Str) Str {
	return Words(s)[0]
}

func LastWord(s Str) Str {
	return LastElem(Words(s))
}

func IsEmptyStripStr(s Str) bool {
	return Strip(s) == ""
}

func PanicIfEmptyStripStr(s Str) {
	if IsEmptyStripStr(s) {
		Panic("string is empty")
	}
}

func StripOrPanicIfEmpty(s Str) Str {
	s = Strip(s)
	PanicAnyNotInit(s)
	return s
}
