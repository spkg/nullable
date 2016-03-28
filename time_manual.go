package nullable

// This file is manually prepared, because time.Time
// works a bit differently to the other types for Scan.

import (
	"fmt"
	"reflect"
	"time"
)

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
