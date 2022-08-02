package golangCommonFuncs

func Swap[T any](value1, value2 *T) {
	*value1, *value2 = *value2, *value1
}
