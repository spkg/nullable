package nullable

// Do not modify. Generated by nullable-generate.

import (
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt64(t *testing.T) {
	testCases := []struct {
		ScanValue     interface{}
		ExpectedError bool
		ExpectedValid bool
		ExpectedValue int64
		JSONText      string
	}{
		{
			ScanValue:     int64(11),
			ExpectedValid: true,
			ExpectedValue: 11,
			JSONText:      `11`,
		},
		{
			ScanValue:     uint64(12),
			ExpectedValid: true,
			ExpectedValue: 12,
			JSONText:      `12`,
		},
		{
			ScanValue:     int32(13),
			ExpectedValid: true,
			ExpectedValue: 13,
			JSONText:      `13`,
		},
		{
			ScanValue:     uint32(14),
			ExpectedValid: true,
			ExpectedValue: 14,
			JSONText:      `14`,
		},
		{
			ScanValue:     []byte("string value"),
			ExpectedError: true,
			ExpectedValid: false,
			ExpectedValue: 0,
			JSONText:      `null`,
		},
		{
			ScanValue:     nil,
			ExpectedValid: false,
			ExpectedValue: 0,
			JSONText:      "null",
		},
	}
	assert := assert.New(t)
	for i, tc := range testCases {
		tcName := fmt.Sprintf("test case %d", i)
		var nv Int64
		err := nv.Scan(tc.ScanValue)
		if tc.ExpectedError {
			assert.Error(err, tcName)
			continue
		} else {
			assert.NoError(err, tcName)
			assert.Equal(tc.ExpectedValid, nv.Valid, tcName)
			assert.Equal(tc.ExpectedValue, nv.Int64, tcName)
		}
		v, err := nv.Value()
		assert.NoError(err)
		if tc.ExpectedValid {
			assert.Equal(driver.Value(int64(tc.ExpectedValue)), v, tcName)
			assert.NotNil(nv.Ptr(), tcName)
			assert.Equal(nv.Int64, *(nv.Ptr()), tcName)
			nv2 := Int64FromPtr(nv.Ptr())
			assert.Equal(nv, nv2, tcName)
		} else {
			assert.Nil(v, tcName)
			assert.Nil(nv.Ptr(), tcName)
			nv2 := Int64FromPtr(nv.Ptr())
			assert.Equal(nv, nv2, tcName)
		}
		jsonText, err := nv.MarshalJSON()
		assert.NoError(err)
		assert.Equal(tc.JSONText, string(jsonText), tcName)
		var nt2 Int64
		err = nt2.UnmarshalJSON(jsonText)
		assert.NoError(err)
		assert.Equal(nv.Valid, nt2.Valid, tcName)
		// invalid JSON for any type
		err = nt2.UnmarshalJSON([]byte("00 this is not valid xx"))
		assert.Error(err)

		// test Normalized comparison
		{
			n1 := Int64{
				Int64: 1,
			}

			n2 := Int64{
				Int64: 0,
			}

			n3 := Int64{
				Int64: 1,
				Valid: true,
			}

			if n1.Normalized() != n2.Normalized() {
				t.Errorf("expected equal, got not equal: %v != %v", n1, n2)
			}
			if n3.Normalized() != n3.Normalized() {
				t.Errorf("expected equal, got not equal: %v != %v", n3, n3)
			}
		}
	}
}
