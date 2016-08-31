package nullable

// Do not modify. Generated by nullable-generate.

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Uint represents a uint value that may be null.
// This type implements the Scanner interface so it
// can be used as a scan destination, similar to NullString.
// It also implements the necessary interfaces to serialize
// to and from JSON.
type Uint struct {
	Uint  uint
	Valid bool
}

// UintFromPtr returns a Uint whose value matches ptr.
func UintFromPtr(ptr *uint) Uint {
	var v Uint
	return v.Assign(ptr)
}

// Assign the value of the pointer. If the pointer is nil,
// then then Valid is false, otherwise Valid is true.
func (n *Uint) Assign(ptr *uint) Uint {
	if ptr == nil {
		n.Valid = false
		n.Uint = 0
	} else {
		n.Valid = true
		n.Uint = *ptr
	}
	return *n
}

// Ptr returns a pointer to uint. If Valid is false
// then the pointer is nil, otherwise it is non-nil.
func (n Uint) Ptr() *uint {
	if n.Valid {
		v := n.Uint
		return &v
	}
	return nil
}

// Normalized returns a Uint that can be compared with
// another Uint for equality.
func (n Uint) Normalized() Uint {
	if n.Valid {
		return n
	}
	// If !Valid, then Uint could be any value.
	// Normalized value can be compared for equality.
	return Uint{}
}

// Scan implements the sql.Scanner interface.
func (n *Uint) Scan(value interface{}) error {
	var nt sql.NullInt64
	err := nt.Scan(value)
	if err != nil {
		return err
	}
	n.Valid = nt.Valid
	n.Uint = uint(nt.Int64)
	return nil
}

// Value implements the driver.Valuer interface.
func (n Uint) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Uint), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (n Uint) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Uint)
	}
	return []byte("null"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Uint) UnmarshalJSON(p []byte) error {
	if bytes.Equal(p, jsonNull) {
		n.Uint = 0
		n.Valid = false
		return nil
	}

	var v uint
	if err := json.Unmarshal(p, &v); err != nil {
		return err
	}

	n.Uint = v
	n.Valid = true
	return nil
}
