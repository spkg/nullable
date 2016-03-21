package nullable

// Do not modify. Generated by nullable-generate.

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Uint16 represents a uint16 value that may be null.
// This type implements the Scanner interface so it
// can be used as a scan destination, similar to NullString.
// It also implements the necessary interfaces to serialize
// to and from JSON.
type Uint16 struct {
	Uint16 uint16
	Valid  bool
}

// Assign the value of the pointer. If the pointer is nil,
// then then Valid is false, otherwise Valid is true.
func (n *Uint16) Assign(ptr *uint16) Uint16 {
	if ptr == nil {
		n.Valid = false
		n.Uint16 = 0
	} else {
		n.Valid = true
		n.Uint16 = *ptr
	}
	return *n
}

// Pointer returns a pointer to uint16. If Valid is false
// then the pointer is nil, otherwise it is non-nil.
func (n Uint16) Pointer() *uint16 {
	if n.Valid {
		v := n.Uint16
		return &v
	}
	return nil
}

// Scan implements the sql.Scanner interface.
func (n *Uint16) Scan(value interface{}) error {
	var nt sql.NullInt64
	err := nt.Scan(value)
	if err != nil {
		return err
	}
	n.Valid = nt.Valid
	n.Uint16 = uint16(nt.Int64)
	return nil
}

// Value implements the driver.Valuer interface.
func (n Uint16) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Uint16), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (n Uint16) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Uint16)
	}
	return []byte("null"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Uint16) UnmarshalJSON(p []byte) error {
	if bytes.Equal(p, []byte("null")) {
		n.Uint16 = 0
		n.Valid = false
		return nil
	}

	var v uint16
	if err := json.Unmarshal(p, &v); err != nil {
		return err
	}

	n.Uint16 = v
	n.Valid = true
	return nil
}
