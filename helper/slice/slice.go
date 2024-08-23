package slice

func ToAnySlice[T any](v []T) []any {
	result := make([]any, len(v))
	for i := range v {
		result[i] = v[i]
	}
	return result
}
