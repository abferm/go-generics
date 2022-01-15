package generics

// Zero returns the zero value for the specified type
func Zero[T any]() T {
	var z T
	return z
}
