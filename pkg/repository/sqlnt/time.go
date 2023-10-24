package sqlnt

import (
	"database/sql"
	"time"
)

func ToTimeNull(value *time.Time) sql.NullTime {
	var nullTime sql.NullTime

	if value != nil {
		nullTime.Time = *value
		nullTime.Valid = true
	}

	return nullTime
}

func ToTime(value *sql.NullTime) *time.Time {
	if value.Valid {
		return &value.Time
	}
	return nil
}
