package gofuncs

import (
	"strconv"
	"time"
)

func ToEmptyInterface[T BasicType](value T) any {
	var emptyInterface interface{} = value
	return emptyInterface
}

func ConvertToAnyTypeSeq[T BasicType](values ...T) []any {
	var convertedSeq []any
	for _, value := range values {
		convertedSeq = append(convertedSeq, ToEmptyInterface(value))
	}
	return convertedSeq
}

func StrToBool(value string) bool {
	res, err := strconv.ParseBool(value)
	CheckErr(err)
	return res
}

func StrToInt(value string) int {
	res, err := strconv.Atoi(value)
	CheckErr(err)
	return res
}

func StrToIntToFloat(value string) Float {
	return Float(StrToInt(value))
}

func NumberToPct(value int) Float {
	return float64(value) / 100
}

func StrToPct(value string) Float {
	return StrToIntToFloat(value) / 100
}

func StrToFloat(value string) Float {
	res, err := strconv.ParseFloat(value, 64)
	CheckErr(err)
	return res
}

func NumberToMillis[T Number](value T) time.Duration {
	return time.Duration(Float(value)*1000) * time.Microsecond
}

func StrToMillis(value string) time.Duration {
	number := StrToFloat(value)
	return NumberToMillis(number)
}
