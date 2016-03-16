package nullable

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	testCases := []struct {
		ScanValue     interface{}
		ExpectedError string
		ExpectedValid bool
		ExpectedValue time.Time
		JSONText      string
	}{
		{
			ScanValue:     time.Date(2001, 11, 10, 15, 04, 05, 0, time.FixedZone("AEST", 10*3600)),
			ExpectedError: "",
			ExpectedValid: true,
			ExpectedValue: time.Date(2001, 11, 10, 15, 04, 05, 0, time.FixedZone("AEST", 10*3600)),
			JSONText:      `"2001-11-10T15:04:05+10:00"`,
		},
		{
			ScanValue:     53.5,
			ExpectedError: "cannot convert float64 to time",
			ExpectedValid: false,
			ExpectedValue: time.Time{},
			JSONText:      `null`,
		},
		{
			ScanValue:     nil,
			ExpectedError: "",
			ExpectedValid: false,
			ExpectedValue: time.Time{},
			JSONText:      "null",
		},
	}
	assert := assert.New(t)
	for i, tc := range testCases {
		tcName := fmt.Sprintf("test case %d", i)
		var nv Time
		err := nv.Scan(tc.ScanValue)
		if tc.ExpectedError != "" {
			assert.Error(err)
			assert.True(strings.Contains(err.Error(), tc.ExpectedError), tcName)
		} else {
			assert.NoError(err, tcName)
			assert.Equal(tc.ExpectedValid, nv.Valid)
			assert.Equal(tc.ExpectedValue, nv.Time)
		}
		v, err := nv.Value()
		assert.NoError(err)
		if tc.ExpectedValid {
			assert.Equal(driver.Value(tc.ExpectedValue), v)
		} else {
			assert.Nil(v)
		}
		jsonText, err := nv.MarshalJSON()
		assert.NoError(err)
		assert.Equal(tc.JSONText, string(jsonText))
		var nt2 Time
		err = nt2.UnmarshalJSON(jsonText)
		assert.NoError(err)
		assert.Equal(nv.Valid, nt2.Valid)
		assert.True(nv.Time.Equal(nt2.Time))
	}

	// attempt to unmarshal invalid format
	var nt3 Time
	err := nt3.UnmarshalJSON([]byte("xxxyyy"))
	assert.Error(err)
}
