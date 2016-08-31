package nullable

// Do not modify. Generated by nullable-generate.

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Int16 represents an int16 value that may be null.
// This type implements the Scanner interface so it
// can be used as a scan destination, similar to NullString.
// It also implements the necessary interfaces to serialize
// to and from JSON.
type Int16 struct {
	Int16 int16
	Valid bool
}

// Int16FromPtr returns a Int16 whose value matches ptr.
func Int16FromPtr(ptr *int16) Int16 {
	var v Int16
	return v.Assign(ptr)
}

// Assign the value of the pointer. If the pointer is nil,
// then then Valid is false, otherwise Valid is true.
func (n *Int16) Assign(ptr *int16) Int16 {
	if ptr == nil {
		n.Valid = false
		n.Int16 = 0
	} else {
		n.Valid = true
		n.Int16 = *ptr
	}
	return *n
}

// Ptr returns a pointer to int16. If Valid is false
// then the pointer is nil, otherwise it is non-nil.
func (n Int16) Ptr() *int16 {
	if n.Valid {
		v := n.Int16
		return &v
	}
	return nil
}

// Normalized returns an Int16 that can be compared with
// another Int16 for equality.
func (n Int16) Normalized() Int16 {
	if n.Valid {
		return n
	}
	// If !Valid, then Int16 could be any value.
	// Normalized value can be compared for equality.
	return Int16{}
}

// Scan implements the sql.Scanner interface.
func (n *Int16) Scan(value interface{}) error {
	var nt sql.NullInt64
	err := nt.Scan(value)
	if err != nil {
		return err
	}
	n.Valid = nt.Valid
	n.Int16 = int16(nt.Int64)
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
	return []byte("null"), nil
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
