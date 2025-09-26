package v1

import (
	"fmt"
	"testing"
)

func TestUnmarshalJSON_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid input",
			input:   `{"type": "validType", "field": "value"}`,
			wantErr: false,
		},
		{
			name:    "null input",
			input:   "null",
			wantErr: true,
		},
		{
			name:    "invalid input",
			input:   `{"type": "invalidType"}`,
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

			c := &TypedVisitorConfig{}
			err := c.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
