package ref

func ToPtr[T any](s T) *T {
	return &s
}

func Deref[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}
