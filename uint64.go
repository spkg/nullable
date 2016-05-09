package nullable

// Do not modify. Generated by nullable-generate.

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Uint64 represents a uint64 value that may be null.
// This type implements the Scanner interface so it
// can be used as a scan destination, similar to NullString.
// It also implements the necessary interfaces to serialize
// to and from JSON.
type Uint64 struct {
	Uint64 uint64
	Valid  bool
}

// Uint64FromPtr returns a Uint64 whose value matches ptr.
func Uint64FromPtr(ptr *uint64) Uint64 {
	var v Uint64
	return v.Assign(ptr)
}

// Assign the value of the pointer. If the pointer is nil,
// then then Valid is false, otherwise Valid is true.
func (n *Uint64) Assign(ptr *uint64) Uint64 {
	if ptr == nil {
		n.Valid = false
		n.Uint64 = 0
	} else {
		n.Valid = true
		n.Uint64 = *ptr
	}
	return *n
}

// Ptr returns a pointer to uint64. If Valid is false
// then the pointer is nil, otherwise it is non-nil.
func (n Uint64) Ptr() *uint64 {
	if n.Valid {
		v := n.Uint64
		return &v
	}
	return nil
}

// Scan implements the sql.Scanner interface.
func (n *Uint64) Scan(value interface{}) error {
	var nt sql.NullInt64
	err := nt.Scan(value)
	if err != nil {
		return err
	}
	n.Valid = nt.Valid
	n.Uint64 = uint64(nt.Int64)
	return nil
}

// Value implements the driver.Valuer interface.
func (n Uint64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Uint64), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (n Uint64) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Uint64)
	}
	return []byte("null"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Uint64) UnmarshalJSON(p []byte) error {
	if bytes.Equal(p, jsonNull) {
		n.Uint64 = 0
		n.Valid = false
		return nil
	}

	var v uint64
	if err := json.Unmarshal(p, &v); err != nil {
		return err
	}

	n.Uint64 = v
	n.Valid = true
	return nil
}
