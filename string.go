package nullable

// Do not modify. Generated by nullable-generate.

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// String represents a string value that may be null.
// This type implements the Scanner interface so it
// can be used as a scan destination, similar to NullString.
// It also implements the necessary interfaces to serialize
// to and from JSON.
type String struct {
	String string
	Valid  bool
}

// StringFromPtr returns a String whose value matches ptr.
func StringFromPtr(ptr *string) String {
	var v String
	return v.Assign(ptr)
}

// Assign the value of the pointer. If the pointer is nil,
// then then Valid is false, otherwise Valid is true.
func (s *String) Assign(ptr *string) String {
	if ptr == nil {
		s.Valid = false
		s.String = ""
	} else {
		s.Valid = true
		s.String = *ptr
	}
	return *s
}

// Ptr returns a pointer to string. If Valid is false
// then the pointer is nil, otherwise it is non-nil.
func (s String) Ptr() *string {
	if s.Valid {
		v := s.String
		return &v
	}
	return nil
}

// Scan implements the sql.Scanner interface.
func (s *String) Scan(value interface{}) error {
	var nt sql.NullString
	err := nt.Scan(value)
	if err != nil {
		return err
	}
	s.Valid = nt.Valid
	s.String = nt.String

	return nil
}

// Value implements the driver.Valuer interface.
func (s String) Value() (driver.Value, error) {
	if !s.Valid {
		return nil, nil
	}
	return string(s.String), nil
}

// MarshalJSON implements the json.Marshaler interface.
func (s String) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return []byte("null"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *String) UnmarshalJSON(p []byte) error {
	if bytes.Equal(p, []byte("null")) {
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
