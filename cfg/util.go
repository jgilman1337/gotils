package cfg

// Helper function to get the zero value of any type.
func Zero[T any]() T {
	var zero T
	return zero
}
