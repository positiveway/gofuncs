package main

import (
	"github.com/jinzhu/copier"
	"log"
	"math"
	"reflect"
)

type SignedInt interface {
	int | int8 | int16 | int32 | int64
}

type UnsignedInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Int interface {
	SignedInt | UnsignedInt
}

type Float = float64

type FloatNumber interface {
	float32 | float64
}

type SignedNumber interface {
	SignedInt | FloatNumber
}

func PanicUnsupportedType(value any) {
	log.Fatalf("Type \"%s\" is not supported", reflect.TypeOf(value).String())
}

func IsNotInit(value any) bool {
	switch v := value.(type) {
	case float32:
		return math.IsNaN(float64(v))
	case float64:
		return math.IsNaN(v)
	case string:
		return v == ""
	default:
		PanicUnsupportedType(v)
	}
	return false
}

func Copy[T any](value *T) *T {
	copiedValue := new(T)

	err := copier.Copy(copiedValue, value)
	if err != nil {
		panic("Copying failed")
	}

	return copiedValue
}

type STG string

func main() {
	var g, c *int
	c = new(int)
	*c = 5
	g = Copy(c)
	println(*g)

	var b STG = "0"
	print(IsNotInit(b))

	var a = math.NaN()
	print(a >= 0)
}
