package main

import "math"

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

func main() {
	var a = math.NaN()
	print(a >= 0)
}
