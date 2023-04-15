package funcs

// Contains returns true if one of the given values satisfy the predicate.
func Contains[T comparable, P func(value T) bool](values []T, predicate P) bool {
	for i := range values {
		if predicate(values[i]) {
			return true
		}
	}
	return false
}

// Map transform the given values into another slice, and map each value to another value by using the given mapper.
func Map[I, O any](values []I, mapper func(inp I) O) []O {
	return Reduce(values, make([]O, 0, len(values)), func(acc []O, cur I) []O {
		return append(acc, mapper(cur))
	})
}

// Reduce executes a user-supplied "reducer" callback function on each element of the array, in order,
// passing in the return value from the calculation on the preceding element.
func Reduce[I, O any](values []I, init O, reducer func(accumulator O, current I) O) O {
	var accumulator = init
	for _, inp := range values {
		accumulator = reducer(accumulator, inp)
	}
	return accumulator
}

// Values get the values of map.
func Values[K comparable, V any](m map[K]V) []V {
	outs := make([]V, 0, len(m))
	for k := range m {
		outs = append(outs, m[k])
	}
	return outs
}
