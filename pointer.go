package pointer

// To returns pointer to any Go value.
func To[T any](val T) *T {
	return &val
}

// Val returns value of pointer or default value if pointer is 'nil'.
func Val[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}
