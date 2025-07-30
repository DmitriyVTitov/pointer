package pointer

// To returns pointer to any Go value.
func To[T any](val T) *T {
	return &val
}
