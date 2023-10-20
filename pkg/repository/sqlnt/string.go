package sqlnt

import "database/sql"

func ToStringNull(value *string) sql.NullString {
	var nullString sql.NullString

	if value != nil {
		nullString.String = *value
		nullString.Valid = true
	}

	return nullString
}

func ToString(value *sql.NullString) *string {
	if value.Valid {
		return &value.String
	}
	return nil
}
