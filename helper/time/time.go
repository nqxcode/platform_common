package time

import (
	"database/sql"

	"github.com/nqxcode/platform_common/pointer"
)

// ToUnixNanoFromSQLNullTime convert sql null time to unix nano
func ToUnixNanoFromSQLNullTime(value sql.NullTime) *int64 {
	if !value.Valid {
		return nil
	}
	return pointer.ToPtr(value.Time.UnixNano())
}
