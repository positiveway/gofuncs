package golangCommonFuncs

import "math"

const NanUInt = -1

func NaN() Float {
	return math.NaN()
}

func IsPositive[T Number](value T) bool {
	return value > 0
}

func Abs[T SignedNumber](val T) T {
	return T(math.Abs(Float(val)))
}

func IsNegative[T SignedNumber](val T) bool {
	return math.Signbit(Float(val))
}

func SignAsNumber[T SignedNumber](val T) T {
	res := T(1)
	if IsNegative(val) {
		res *= -1
	}
	return res
}

func ApplySign[T SignedNumber](isNegative bool, val T) T {
	if isNegative {
		val *= -1
	}
	return val
}

func GetIsNegativeAndAbs[T SignedNumber](val T) (bool, T) {
	return IsNegative(val), Abs(val)
}

func Trunc[T Number](number Float, precision T) Float {
	multiplier := math.Pow(10, Float(precision))
	return math.Trunc(number*multiplier) / multiplier
}

func Round[T Number](number Float, precision T) Float {
	multiplier := math.Pow(10, Float(precision))
	return math.Round(number*multiplier) / multiplier
}

func FloatToIntRound[T Int](value Float) T {
	return T(math.Round(value))
}

func Sqr[T Number](x T) T {
	return x * x
}

func Sqrt[T Number](val T) Float {
	return math.Sqrt(Float(val))
}
