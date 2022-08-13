package gofuncs

import (
	"strconv"
	"time"
)

func StrToBool(value Str) bool {
	res, err := strconv.ParseBool(value)
	CheckErr(err)
	return res
}

func StrToInt(value Str) int {
	res, err := strconv.Atoi(value)
	CheckErr(err)
	return res
}

func StrToIntToFloat(value Str) Float {
	return Float(StrToInt(value))
}

func NumberToPct(value int) Float {
	return float64(value) / 100
}

func StrToPct(value Str) Float {
	return StrToIntToFloat(value) / 100
}

func StrToFloat(value Str) Float {
	res, err := strconv.ParseFloat(value, 64)
	CheckErr(err)
	return res
}

func NumberToMillis[T Number](value T) time.Duration {
	return time.Duration(Float(value)*1000) * time.Microsecond
}

func StrToMillis(value Str) time.Duration {
	number := StrToFloat(value)
	return NumberToMillis(number)
}
