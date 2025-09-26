package models

import (
	"fmt"
	"testing"
)

func TestSetRaw_7(t *testing.T) {
	testCases := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{
			name:    "uint64 value",
			value:   uint64(18446744073709551615),
			wantErr: false,
		},
		{
			name:    "string value",
			value:   "18446744073709551615",
			wantErr: false,
		},
		{
			name:    "invalid string value",
			value:   "invalid",
			wantErr: true,
		},
		{
			name:    "unsupported type",
			value:   []byte("unsupported"),
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var field PositiveBigIntegerField
			err := field.SetRaw(tc.value)
			if (err != nil) != tc.wantErr {
				t.Errorf("SetRaw() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
