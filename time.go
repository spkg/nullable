package nullable

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// Time is a nullable string.
type Time struct {
	Time  time.Time
	Valid bool
}

// Scan implements the sql.Scanner interface.
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		t.Time, t.Valid = time.Time{}, false
		return nil
	}
	if tm, ok := value.(time.Time); ok {
		t.Valid = true
		t.Time = tm
		return nil
	}
	typ := reflect.TypeOf(value)
	return fmt.Errorf("cannot convert %s to time", typ.Name())
}

// Value implements the driver.Valuer interface.
func (t Time) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

// MarshalJSON implements the json.Marshaler interface.
func (t Time) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	}
	return jsonNull, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Time) UnmarshalJSON(p []byte) error {
	if bytes.Equal(p, jsonNull) {
		t.Time = time.Time{}
		t.Valid = false
		return nil
	}

	var v time.Time
	if err := json.Unmarshal(p, &v); err != nil {
		return err
	}

	t.Time = v
	t.Valid = true
	return nil
}
