package nullable

// Do not modify. Generated by nullable-generate.

import (
	"bytes"

	"database/sql/driver"
	"encoding/json"
	"time"
)

// Time represents a time.Time value that may be null.
// This type implements the Scanner interface so it
// can be used as a scan destination, similar to NullString.
// It also implements the necessary interfaces to serialize
// to and from JSON.
type Time struct {
	Time  time.Time
	Valid bool
}

// TimeFromPtr returns a Time whose value matches ptr.
func TimeFromPtr(ptr *time.Time) Time {
	var v Time
	return v.Assign(ptr)
}

// Assign the value of the pointer. If the pointer is nil,
// then then Valid is false, otherwise Valid is true.
func (tm *Time) Assign(ptr *time.Time) Time {
	if ptr == nil {
		tm.Valid = false
		tm.Time = time.Time{}
	} else {
		tm.Valid = true
		tm.Time = *ptr
	}
	return *tm
}

// Ptr returns a pointer to time.Time. If Valid is false
// then the pointer is nil, otherwise it is non-nil.
func (tm Time) Ptr() *time.Time {
	if tm.Valid {
		v := tm.Time
		return &v
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (tm Time) Value() (driver.Value, error) {
	if !tm.Valid {
		return nil, nil
	}
	return tm.Time, nil

}

// MarshalJSON implements the json.Marshaler interface.
func (tm Time) MarshalJSON() ([]byte, error) {
	if tm.Valid {
		return json.Marshal(tm.Time)
	}
	return []byte("null"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (tm *Time) UnmarshalJSON(p []byte) error {
	if bytes.Equal(p, []byte("null")) {
		tm.Time = time.Time{}
		tm.Valid = false
		return nil
	}

	var v time.Time
	if err := json.Unmarshal(p, &v); err != nil {
		return err
	}

	tm.Time = v
	tm.Valid = true
	return nil
}
