package gofuncs

import "math"

func IsNotInit(value any) bool {
	switch v := value.(type) {
	case Float:
		return math.IsNaN(v)
	case Str:
		return v == ""
	case int, int32, int64:
		return v == NanUInt
	default:
		Panic("Type is not supported")
	}
	return false
}

func AnyNotInit(values ...any) bool {
	_, found := IsAnyPredicate(values, IsNotInit)
	return found
}

func PanicNotInit() {
	Panic("Value is not initialized")
}

func PanicAnyNotInit(values ...any) {
	if AnyNotInit(values...) {
		PanicNotInit()
	}
}

func IsAnyPredicate[T any](values []T, predicate func(value T) bool) (T, bool) {
	for _, value := range values {
		if predicate(value) {
			return value, true
		}
	}
	return nil, false
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
