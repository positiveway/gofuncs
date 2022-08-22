package gofuncs

import (
	"math"
	"reflect"
)

func PanicUnsupportedType(value any) {
	Panic("Type \"%s\" is not supported", reflect.TypeOf(value).String())
}

func IsNotInit(value any) bool {
	switch v := value.(type) {
	case float32:
		return math.IsNaN(float64(v))
	case float64:
		return math.IsNaN(v)
	case Str:
		return v == ""
	default:
		PanicUnsupportedType(v)
	}
	return false
}

func IsEmpty(value any) bool {
	switch v := value.(type) {
	case float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v == 0
	case Str:
		return IsEmptyStripStr(v)
	default:
		PanicUnsupportedType(v)
	}
	return false
}

func IsNotInitOrEmpty(value any) bool {
	switch v := value.(type) {
	case float32, float64:
		return IsNotInit(v) || IsEmpty(v)
	default:
		return IsEmpty(v)
	}
}

func AnyNotInit(values ...any) bool {
	_, found := IsAnyPredicate(values, IsNotInit)
	return found
}

func AnyIsEmpty(values ...any) bool {
	_, found := IsAnyPredicate(values, IsEmpty)
	return found
}

func AnyNotInitOrEmpty(values ...any) bool {
	_, found := IsAnyPredicate(values, IsNotInitOrEmpty)
	return found
}

func PanicNotInit() {
	Panic("Value is not initialized")
}

func PanicIsEmpty() {
	Panic("Value is empty")
}

func PanicNotInitOrEmpty() {
	Panic("Value is not initialized or is empty")
}

func PanicAnyNotInit(values ...any) {
	if AnyNotInit(values...) {
		PanicNotInit()
	}
}

func PanicAnyIsEmpty(values ...any) {
	if AnyIsEmpty(values...) {
		PanicIsEmpty()
	}
}

func PanicAnyNotInitOrEmpty(values ...any) {
	if AnyNotInitOrEmpty(values...) {
		PanicNotInitOrEmpty()
	}
}

func IsAnyPredicate[T any](values []T, predicate func(value T) bool) (T, bool) {
	for _, value := range values {
		if predicate(value) {
			return value, true
		}
	}
	var emptyRes T
	return emptyRes, false
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
