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

type Float interface {
	float32 | float64
}

type SignedNumber interface {
	SignedInt | Float
}

type Number interface {
	Int | Float
}

type BasicType interface {
	Number | string | bool | rune
}

func PanicUnsupportedType(value any) {
	log.Fatalf("Type \"%s\" is not supported", reflect.TypeOf(value).String())
}

type STG string

type Address struct {
	Street      string
	HouseNumber int
}

type Person struct {
	Address *Address
	Name    string
}

type University struct {
	Name    string
	Student Person
}

func ToEmptyInterface[T any](value T) any {
	var emptyInterface interface{} = value
	return emptyInterface
}

func main() {

	newUni := University{Student: Person{Address: &Address{Street: "Traidmill"}}}

	nextUni := newUni
	println(nextUni.Student.Address.Street)
	newUni.Student.Address.Street = "ABC"
	println(newUni.Student.Address.Street)
	println(nextUni.Student.Address.Street)

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
