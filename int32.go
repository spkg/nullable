package nullable

// Do not modify. Generated by nullable-generate.

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Int32 represents an int32 value that may be null.
// This type implements the Scanner interface so it
// can be used as a scan destination, similar to NullString.
// It also implements the necessary interfaces to serialize
// to and from JSON.
type Int32 struct {
	Int32 int32
	Valid bool
}

// Assign the value of the pointer. If the pointer is nil,
// then then Valid is false, otherwise Valid is true.
func (n *Int32) Assign(ptr *int32) Int32 {
	if ptr == nil {
		n.Valid = false
		n.Int32 = 0
	} else {
		n.Valid = true
		n.Int32 = *ptr
	}
	return *n
}

// Pointer returns a pointer to int32. If Valid is false
// then the pointer is nil, otherwise it is non-nil.
func (n Int32) Pointer() *int32 {
	if n.Valid {
		v := n.Int32
		return &v
	}
	return nil
}

// Scan implements the sql.Scanner interface.
func (n *Int32) Scan(value interface{}) error {
	var nt sql.NullInt64
	err := nt.Scan(value)
	if err != nil {
		return err
	}
	n.Valid = nt.Valid
	n.Int32 = int32(nt.Int64)
	return nil
}

// Value implements the driver.Valuer interface.
func (n Int32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Int32), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (n Int32) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Int32)
	}
	return []byte("null"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Int32) UnmarshalJSON(p []byte) error {
	if bytes.Equal(p, []byte("null")) {
		n.Int32 = 0
		n.Valid = false
		return nil
	}

	var v int32
	if err := json.Unmarshal(p, &v); err != nil {
		return err
	}

	n.Int32 = v
	n.Valid = true
	return nil
}