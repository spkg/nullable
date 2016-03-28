package nullable

// Do not modify. Generated by nullable-generate.

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	testCases := []struct {
		ScanValue     interface{}
		ExpectedError string
		ExpectedValid bool
		ExpectedValue string
		JSONText      string
	}{
		{
			ScanValue:     "string-val",
			ExpectedError: "",
			ExpectedValid: true,
			ExpectedValue: "string-val",
			JSONText:      `"string-val"`,
		},
		{
			ScanValue:     []byte("bytes"),
			ExpectedError: "",
			ExpectedValid: true,
			ExpectedValue: "bytes",
			JSONText:      `"bytes"`,
		},
		{
			ScanValue:     nil,
			ExpectedError: "",
			ExpectedValid: false,
			ExpectedValue: "",
			JSONText:      "null",
		},
		{
			ScanValue:     int64(99),
			ExpectedError: "",
			ExpectedValid: true,
			ExpectedValue: "99",
			JSONText:      "\"99\"",
		},
		{
			ScanValue:     false,
			ExpectedError: "",
			ExpectedValid: true,
			ExpectedValue: "false",
			JSONText:      "\"false\"",
		},
	}
	assert := assert.New(t)
	for i, tc := range testCases {
		tcName := fmt.Sprintf("test case %d", i)
		var nv String
		err := nv.Scan(tc.ScanValue)
		if tc.ExpectedError != "" {
			assert.Error(err, tcName)
			assert.True(strings.Contains(err.Error(), tc.ExpectedError), err.Error())
			continue
		} else {
			assert.NoError(err, tcName)
			assert.Equal(tc.ExpectedValid, nv.Valid, tcName)
			assert.Equal(tc.ExpectedValue, nv.String, tcName)
		}
		v, err := nv.Value()
		assert.NoError(err)
		if tc.ExpectedValid {
			assert.Equal(driver.Value(string(tc.ExpectedValue)), v, tcName)
			assert.NotNil(nv.Ptr(), tcName)
			assert.Equal(nv.String, *(nv.Ptr()), tcName)
			nv2 := StringFromPtr(nv.Ptr())
			assert.Equal(nv, nv2, tcName)
		} else {
			assert.Nil(v, tcName)
			assert.Nil(nv.Ptr(), tcName)
			nv2 := StringFromPtr(nv.Ptr())
			assert.Equal(nv, nv2, tcName)
		}
		jsonText, err := nv.MarshalJSON()
		assert.NoError(err)
		assert.Equal(tc.JSONText, string(jsonText), tcName)
		var nt2 String
		err = nt2.UnmarshalJSON(jsonText)
		assert.NoError(err)
		assert.Equal(nv.Valid, nt2.Valid, tcName)
		// invalid JSON for any type
		err = nt2.UnmarshalJSON([]byte("00 this is not valid xx"))
		assert.Error(err)
	}
}
