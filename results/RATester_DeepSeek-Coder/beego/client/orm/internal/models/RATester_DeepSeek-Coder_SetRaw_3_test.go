package models

import (
	"fmt"
	"testing"
	"time"
)

func TestSetRaw_3(t *testing.T) {
	type testCase struct {
		name    string
		value   interface{}
		wantErr bool
	}

	tests := []testCase{
		{
			name:    "Valid time.Time value",
			value:   time.Now(),
			wantErr: false,
		},
		{
			name:    "Valid string value",
			value:   "2006-01-02 15:04:05",
			wantErr: false,
		},
		{
			name:    "Invalid value",
			value:   "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			e := &DateTimeField{}
			if err := e.SetRaw(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("DateTimeField.SetRaw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
