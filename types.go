package gofuncs

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
