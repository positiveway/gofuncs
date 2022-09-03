package main

import (
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

type STG string

func main() {
	var c *int
	c = new(int)
	*c = 5

	var b STG = "0"
	println(b)

	var a = math.NaN()
	println(a >= 0)

	g := 5.5
	println(math.Mod(g, 1) != 0)
}
