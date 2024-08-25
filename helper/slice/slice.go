package slice

import (
	"strconv"

	"github.com/nqxcode/platform_common/pagination"
)

// ToAnySlice typed slice to any slice
func ToAnySlice[T any](v []T) []any {
	result := make([]any, len(v))
	for i := range v {
		result[i] = v[i]
	}
	return result
}

// ToStringSliceFromIntSlice transform int64 slice to string slice
func ToStringSliceFromIntSlice(v []int64) []string {
	result := make([]string, len(v))
	for i := range v {
		result[i] = strconv.FormatInt(v[i], 10)
	}
	return result
}

// ByLimit get slice by limit
func ByLimit[T comparable](v []T, limit pagination.Limit) []T {
	total := uint64(len(v))
	if total == 0 {
		return nil
	}

	offset := limit.Offset
	if offset > total {
		offset = total
	}

	end := limit.Offset + limit.Limit
	if end == 0 || end > total {
		end = total
	}

	return v[offset:end]
}
