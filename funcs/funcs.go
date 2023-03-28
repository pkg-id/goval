package funcs

func Contains[T comparable, P func(value T) bool](values []T, predicate P) bool {
	for i := range values {
		if predicate(values[i]) {
			return true
		}
	}
	return false
}
