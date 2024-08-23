package time

import (
	"database/sql"

	"github.com/nqxcode/platform_common/pointer"
)

func ToUnixNanoFromSqlNullTime(value sql.NullTime) *int64 {
	if !value.Valid {
		return nil
	}
	return pointer.ToPtr(value.Time.UnixNano())
}
