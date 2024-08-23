package slice

import "strconv"

func ToAnySlice[T any](v []T) []any {
	result := make([]any, len(v))
	for i := range v {
		result[i] = v[i]
	}
	return result
}

func ToStringSliceFromIntSlice(v []int64) []string {
	result := make([]string, len(v))
	for i := range v {
		result[i] = strconv.FormatInt(v[i], 10)
	}
	return result
}
