package nullable

import (
	"database/sql/driver"
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
			ScanValue:     "string value 1",
			ExpectedError: "",
			ExpectedValid: true,
			ExpectedValue: "string value 1",
			JSONText:      `"string value 1"`,
		},
		{
			ScanValue:     []byte("string value 2"),
			ExpectedError: "",
			ExpectedValid: true,
			ExpectedValue: "string value 2",
			JSONText:      `"string value 2"`,
		},
		{
			ScanValue:     nil,
			ExpectedError: "",
			ExpectedValid: false,
			ExpectedValue: "",
			JSONText:      "null",
		},
	}
	assert := assert.New(t)
	for _, tc := range testCases {
		var nv String
		err := nv.Scan(tc.ScanValue)
		if tc.ExpectedError != "" {
			assert.Error(err)
			assert.True(strings.Contains(err.Error(), tc.ExpectedError), err.Error())
		} else {
			assert.NoError(err)
			assert.Equal(tc.ExpectedValid, nv.Valid)
			assert.Equal(tc.ExpectedValue, nv.String)
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
		var nt2 String
		err = nt2.UnmarshalJSON(jsonText)
		assert.NoError(err)
		assert.Equal(nv.Valid, nt2.Valid)
		assert.True(nv.String == nt2.String)
	}

	// attempt to unmarshal invalid format
	var nt3 String
	err := nt3.UnmarshalJSON([]byte("xxxyyy"))
	assert.Error(err)
}
