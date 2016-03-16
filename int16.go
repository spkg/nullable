package nullable

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Int16 is a nullable int16.
type Int16 struct {
	Int16 int16
	Valid bool
}

// Scan implements the sql.Scanner interface.
func (n *Int16) Scan(value interface{}) error {
	var ni sql.NullInt64
	err := ni.Scan(value)
	if err != nil {
		return err
	}
	n.Valid = ni.Valid
	n.Int16 = int16(ni.Int64)
	return nil
}

// Value implements the driver.Valuer interface.
func (n Int16) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Int16), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (n Int16) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Int16)
	}
	return jsonNull, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Int16) UnmarshalJSON(p []byte) error {
	if bytes.Equal(p, jsonNull) {
		n.Int16 = 0
		n.Valid = false
		return nil
	}

	var v int16
	if err := json.Unmarshal(p, &v); err != nil {
		return err
	}

	n.Int16 = v
	n.Valid = true
	return nil
}
