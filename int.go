package nullable

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Int is a nullable integer.
type Int struct {
	Int   int
	Valid bool
}

// Scan implements the sql.Scanner interface.
func (n *Int) Scan(value interface{}) error {
	var ni sql.NullInt64
	err := ni.Scan(value)
	if err != nil {
		return err
	}
	n.Valid = ni.Valid
	n.Int = int(ni.Int64)
	return nil
}

// Value implements the driver.Valuer interface.
func (n Int) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Int), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (n Int) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Int)
	}
	return jsonNull, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Int) UnmarshalJSON(p []byte) error {
	if bytes.Equal(p, jsonNull) {
		n.Int = 0
		n.Valid = false
		return nil
	}

	var v int
	if err := json.Unmarshal(p, &v); err != nil {
		return err
	}

	n.Int = v
	n.Valid = true
	return nil
}
