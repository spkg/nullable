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

// Assign the value of the pointer. If the pointer is nil,
// then then Valid is false, otherwise Valid is true.
func (n *Int) Assign(ptr *int) Int {
	if ptr == nil {
		n.Valid = false
		n.Int = 0
	} else {
		n.Valid = true
		n.Int = *ptr
	}
	return *n
}

// Pointer returns a pointer to int. If Valid is false
// then the pointer is nil, otherwise it is non-nil.
func (n Int) Pointer() *int {
	if n.Valid {
		v := n.Int
		return &v
	}
	return nil
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
