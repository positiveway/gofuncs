package gofuncs

import (
	"fmt"
	"strings"
)

func RemoveLastLetters(sourceStr string, numLettersToRemove int) string {
	return sourceStr[:len(sourceStr)-numLettersToRemove]
}

func RemoveFirstLetters(sourceStr string, numLettersToRemove int) string {
	return sourceStr[numLettersToRemove:]
}

func SplitByAnyOf(sourceStr string, separators string) []string {
	if separators == "" {
		Panic("Empty separator")
	}
	var res []string
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

func StartsWith(s string, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func EndsWith(s string, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func StartsWithAnyOf(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if StartsWith(s, prefix) {
			return true
		}
	}
	return false
}

func EndsWithAnyOf(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if EndsWith(s, suffix) {
			return true
		}
	}
	return false
}

func TrimAnyPrefix(s string, prefixes ...string) string {
	for _, prefix := range prefixes {
		if StartsWith(s, prefix) {
			return strings.TrimPrefix(s, prefix)
		}
	}
	return s
}

func Strip(s string) string {
	return strings.TrimSpace(s)
}

func FilterEmptyStrings(slice []string) []string {
	var filtered []string
	for _, s := range slice {
		if !IsEmptyStripStr(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func Split(s string, sep ...string) []string {
	PanicIfEmptyStripStr(s)

	separator := func(separator []string) string {
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

func Words(s string) []string {
	return Split(s)
}

func FirstWord(s string) string {
	return Words(s)[0]
}

func LastWord(s string) string {
	return LastElem(Words(s))
}

func IsEmptyStripStr(s string) bool {
	return Strip(s) == ""
}

func PanicIfEmptyStripStr(s string) {
	if IsEmptyStripStr(s) {
		Panic("string is empty")
	}
}

func StripOrPanicIfEmpty(s string) string {
	s = Strip(s)
	PanicAnyNotInit(s)
	return s
}

func AnyToStr(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func AnyValuesToStr(values ...interface{}) []string {
	var strList []string

	for _, value := range values {
		strList = append(strList, AnyToStr(value))
	}
	return strList
}

func ToLower(value interface{}) string {
	return strings.ToLower(AnyToStr(value))
}

func ToLowerValues(values ...interface{}) []string {
	var loweredValues []string

	for _, value := range values {
		loweredValues = append(loweredValues, ToLower(value))
	}
	return loweredValues
}

func ToLowerAndStrip(value interface{}) string {
	return Strip(ToLower(value))
}

func ToLowerAndStripPanicIfEmpty(value interface{}) string {
	res := ToLowerAndStrip(value)
	PanicIfEmptyStripStr(res)
	return res
}
