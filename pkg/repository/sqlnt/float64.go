package sqlnt

import "database/sql"

func ToF64Null(value *float64) sql.NullFloat64 {
	var nullFloat64 sql.NullFloat64

	if value != nil {
		nullFloat64.Float64 = *value
		nullFloat64.Valid = true
	}

	return nullFloat64
}

func ToF64(value *sql.NullFloat64) *float64 {
	if value.Valid {
		return &value.Float64
	}

	return nil
}
