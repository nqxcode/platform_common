package pointer

// ToPtr to pointer
func ToPtr[T any](s T) *T {
	return &s
}

// Deref dereference of pointer
func Deref[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}
