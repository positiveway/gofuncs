package gofuncs

import (
	"math"
	"reflect"
	"strconv"
	"time"
)

func GetTypeOfEmptyInterface(value any) string {
	return reflect.TypeOf(value).String()
}

func ToEmptyInterface[T any](value T) any {
	var emptyInterface interface{} = value
	return emptyInterface
}

func FromEmptyInterface[T any](value any) T {
	if converted, ok := value.(T); ok {
		return converted
	} else {
		Panic("Conversion of type \"%s\" failed", GetTypeOfEmptyInterface(value))
	}
	panic("")
}

func ConvertToAnyTypeSeq[T any](values ...T) []any {
	var convertedSeq []any
	for _, value := range values {
		convertedSeq = append(convertedSeq, ToEmptyInterface(value))
	}
	return convertedSeq
}

func ConvertFromAnyTypeSeq[T any](values ...any) []T {
	var convertedSeq []T
	for _, value := range values {
		convertedSeq = append(convertedSeq, FromEmptyInterface[T](value))
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

func CheckSourceIsInt(value float64) float64 {
	if math.Mod(value, 1) != 0 {
		Panic("Value is not Integer")
	}
	return value
}

func NumberToPct(value float64) Float {
	PanicAnyNotPositive(value)
	return CheckSourceIsInt(value) / 100
}

func NumberToPctInPlace(value *float64) {
	*value = NumberToPct(*value)
}

func StrToPct(value string) Float {
	return NumberToPct(StrToFloat(value))
}

func StrToFloat(value string) Float {
	res, err := strconv.ParseFloat(value, 64)
	CheckErr(err)
	return res
}

func NumberToMillis[T Number](value T) time.Duration {
	PanicAnyNotPositive(value)
	return time.Duration(Float(value)*1000) * time.Microsecond
}

func StrToMillis(value string) time.Duration {
	number := StrToFloat(value)
	return NumberToMillis(number)
}
