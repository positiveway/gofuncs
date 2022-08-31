package gofuncs

func Pop[K comparable, V any](m map[K]V, key K) V {
	value := m[key]
	delete(m, key)
	return value
}

func AssignWithDuplicateCheck[K comparable, V any](m map[K]V, key K, val V) {
	if _, found := m[key]; found {
		Panic("duplicate position")
	}
	m[key] = val
}

func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	if val, found := m[key]; found {
		return val
	} else {
		return defaultVal
	}
}

func GetOrPanic[K comparable, V any](m map[K]V, key K, msg ...Str) V {
	if val, found := m[key]; found {
		return val
	}
	message := GetPanicMsg(msg, "No such key in map")

	Panic(message+": \"%v\"", key)
	panic("")
}

func IsEmptySlice[T any](seq []T) bool {
	return len(seq) == 0
}

func CheckLengthSlice[T any](seq []T, length int) {
	if !IsPositive(length) {
		Panic("length parameter is incorrect: %v", length)
	}
	if len(seq) < length {
		Panic("Length of sequence should be at least %v", length)
	}
}

func IsEmptyMap[K comparable, V any](seq map[K]V) bool {
	return len(seq) == 0
}

func ShallowCopyMap[K comparable, V any](seq map[K]V) map[K]V {
	copied := map[K]V{}
	for key, val := range seq {
		copied[key] = val
	}
	return copied
}

func LastElem[T any](seq []T) T {
	CheckLengthSlice(seq, 1)
	return seq[len(seq)-1]
}

func Reverse[T BasicType](seq []T) []T {
	var res []T
	for i := len(seq) - 1; i >= 0; i-- {
		res = append(res, seq[i])
	}
	return res
}

func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// A nil argument is equivalent to an empty slice
func IsEqualSlices[T BasicType](a, b []T) bool {
	// this includes nil case
	if IsEmptySlice(a) || IsEmptySlice(b) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func FindInSeqByPredicate[T any](isCurValueBetter func(curValue, prevBestValue T) bool, values ...T) T {
	bestValue := values[0]
	CheckLengthSlice(values, 2)

	for _, value := range values {
		if isCurValueBetter(value, bestValue) {
			bestValue = value
		}
	}
	return bestValue
}

func Max[T Number](values ...T) T {
	return FindInSeqByPredicate(func(curValue, prevBestValue T) bool {
		return curValue > prevBestValue
	}, values...)
}

func Min[T Number](values ...T) T {
	return FindInSeqByPredicate(func(curValue, prevBestValue T) bool {
		return curValue < prevBestValue
	}, values...)
}
