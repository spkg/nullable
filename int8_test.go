package nullable

// Do not modify. Generated by nullable-generate.

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt8(t *testing.T) {
	testCases := []struct {
		ScanValue     interface{}
		ExpectedError string
		ExpectedValid bool
		ExpectedValue int8
		JSONText      string
	}{
		{
			ScanValue:     int64(11),
			ExpectedError: "",
			ExpectedValid: true,
			ExpectedValue: 11,
			JSONText:      `11`,
		},
		{
			ScanValue:     uint64(12),
			ExpectedError: "",
			ExpectedValid: true,
			ExpectedValue: 12,
			JSONText:      `12`,
		},
		{
			ScanValue:     int32(13),
			ExpectedError: "",
			ExpectedValid: true,
			ExpectedValue: 13,
			JSONText:      `13`,
		},
		{
			ScanValue:     uint32(14),
			ExpectedError: "",
			ExpectedValid: true,
			ExpectedValue: 14,
			JSONText:      `14`,
		},
		{
			ScanValue:     []byte("string value"),
			ExpectedError: "invalid syntax",
			ExpectedValid: false,
			ExpectedValue: 0,
			JSONText:      `null`,
		},
		{
			ScanValue:     nil,
			ExpectedError: "",
			ExpectedValid: false,
			ExpectedValue: 0,
			JSONText:      "null",
		},
	}
	assert := assert.New(t)
	for i, tc := range testCases {
		tcName := fmt.Sprintf("test case %d", i)
		var nv Int8
		err := nv.Scan(tc.ScanValue)
		if tc.ExpectedError != "" {
			assert.Error(err, tcName)
			assert.True(strings.Contains(err.Error(), tc.ExpectedError), err.Error())
			continue
		} else {
			assert.NoError(err, tcName)
			assert.Equal(tc.ExpectedValid, nv.Valid)
			assert.Equal(tc.ExpectedValue, nv.Int8)
		}
		v, err := nv.Value()
		assert.NoError(err)
		if tc.ExpectedValid {
			assert.Equal(driver.Value(int64(tc.ExpectedValue)), v)
		} else {
			assert.Nil(v, fmt.Sprintf("test case %d", i))
		}
		jsonText, err := nv.MarshalJSON()
		assert.NoError(err)
		assert.Equal(tc.JSONText, string(jsonText), fmt.Sprintf("test case %d", i))
		var nt2 Int8
		err = nt2.UnmarshalJSON(jsonText)
		assert.NoError(err)
		assert.Equal(nv.Valid, nt2.Valid)
		assert.True(nv.Int8 == nt2.Int8)
	}
}
