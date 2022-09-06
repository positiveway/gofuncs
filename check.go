package gofuncs

import (
	"math"
)

func PanicUnsupportedType(value any) {
	Panic("Type \"%s\" is not supported", GetTypeOfEmptyInterface(value))
}

func GetValueOrDefaultIfEmpty[T BasicType](value, defaultVal T) T {
	if IsEmptyOrNotInit(value) {
		return defaultVal
	} else {
		return value
	}
}

func SetDefaultIfValueIsEmpty[T BasicType](value *T, defaultVal T) {
	*value = GetValueOrDefaultIfEmpty(*value, defaultVal)
}

func IsNotInit(value interface{}) bool {
	switch v := value.(type) {
	case float32:
		return math.IsNaN(float64(v))
	case float64:
		return math.IsNaN(v)
	case string:
		return (v) == ""
	default:
		PanicUnsupportedType(value)
	}
	return false
}

func IsEmpty(value interface{}) bool {
	switch v := value.(type) {
	case float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v == 0
	case string:
		return IsEmptyStripStr(v)
	default:
		PanicUnsupportedType(value)
	}
	return false
}

func IsEmptyOrNotInit(value interface{}) bool {
	switch value.(type) {
	case float32, float64:
		return IsNotInit(value) || IsEmpty(value)
	default:
		return IsEmpty(value)
	}
}

func AnyNotInit[T BasicType](values ...T) bool {
	return IsAnyPredicate(values, IsNotInit[T])
}

func AnyIsEmpty[T BasicType](values ...T) bool {
	return IsAnyPredicate(values, IsEmpty[T])
}

func AnyEmptyOrNotInit[T BasicType](values ...T) bool {
	return IsAnyPredicate(values, IsEmptyOrNotInit[T])
}

//var PanicNotInit = GetPanicWithMsg("Value is not initialized")

func PanicNotInit() {
	Panic("Value is not initialized")
}

func PanicIsEmpty() {
	Panic("Value is empty")
}

func PanicEmptyOrNotInit() {
	Panic("Value is not initialized or is empty")
}

func PanicAnyNotInit[T BasicType](values ...T) {
	PanicIfAny(values, IsNotInit[T], PanicNotInit)
}

func PanicAnyIsEmpty[T BasicType](values ...T) {
	PanicIfAny(values, IsEmpty[T], PanicIsEmpty)
}

func PanicAnyEmptyOrNotInit[T BasicType](values ...T) {
	PanicIfAny(values, IsEmptyOrNotInit[T], PanicEmptyOrNotInit)
}

func GetPanicWithMsg(message string, values ...any) func() {
	return func() {
		Panic(message, values...)
	}
}

func PanicIfAny[T BasicType](values []T, predicate func(value interface{}) bool, panicFunc func()) {
	if IsAnyPredicate(values, predicate) {
		panicFunc()
	}
}

func IsAnyPredicateWithValue[T BasicType](values []T, predicate func(value interface{}) bool) (T, bool) {
	for _, value := range values {
		if predicate(value) {
			return value, true
		}
	}
	var emptyRes T
	return emptyRes, false
}

func IsAnyPredicate[T BasicType](values []T, predicate func(value interface{}) bool) bool {
	_, found := IsAnyPredicateWithValue(values, predicate)
	return found
}

func AnyCmp[T Number](pairs [][]T, cmp func(val1, val2 T) bool) bool {
	for _, pair := range pairs {
		if len(pair) > 2 {
			Panic("Pair can only have 2 elements")
		}
		if cmp(pair[0], pair[1]) {
			return true
		}
	}
	return false
}

func AnyGreater[T Number](pairs [][]T) bool {
	return AnyCmp(pairs, func(val1, val2 T) bool { return val1 > val2 })
}

func AnyGreaterOrEqual[T Number](pairs [][]T) bool {
	return AnyCmp(pairs, func(val1, val2 T) bool { return val1 >= val2 })
}

func AnyLess[T Number](pairs [][]T) bool {
	return AnyCmp(pairs, func(val1, val2 T) bool { return val1 < val2 })
}

func AnyLessOrEqual[T Number](pairs [][]T) bool {
	return AnyCmp(pairs, func(val1, val2 T) bool { return val1 <= val2 })
}

func AnyEqual[T Number](pairs [][]T) bool {
	return AnyCmp(pairs, func(val1, val2 T) bool { return val1 == val2 })
}

func AnyNotEqual[T Number](pairs [][]T) bool {
	return AnyCmp(pairs, func(val1, val2 T) bool { return val1 != val2 })
}

func IsGreaterAbs[T SignedNumber](newValue, oldValue T) bool {
	return Abs(newValue) > Abs(oldValue)
}

type ComparableFields = [][2]any

func IsFieldsEqual(fields ComparableFields) bool {
	for _, pair := range fields {
		if pair[0] != pair[1] {
			return false
		}
	}
	return true
}
