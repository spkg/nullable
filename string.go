package nullable

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// String is a nullable string.
// This type is almost identical to the standard library
// sql.NullString class. It is in this package for completeness only.
type String struct {
	String string
	Valid  bool
}

// Scan implements the sql.Scanner interface.
func (s *String) Scan(value interface{}) error {
	var ns sql.NullString
	err := ns.Scan(value)
	s.Valid = ns.Valid
	s.String = ns.String
	return err
}

// Value implements the driver.Valuer interface.
func (s String) Value() (driver.Value, error) {
	if !s.Valid {
		return nil, nil
	}
	return s.String, nil
}

// MarshalJSON implements the json.Marshaler interface.
func (s String) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return jsonNull, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *String) UnmarshalJSON(p []byte) error {
	if bytes.Equal(p, jsonNull) {
		s.String = ""
		s.Valid = false
		return nil
	}

	var v string
	if err := json.Unmarshal(p, &v); err != nil {
		return err
	}

	s.String = v
	s.Valid = true
	return nil
}
