package types

import (
	"fmt"
	"testing"
)

func TestUnmarshalJSON_2(t *testing.T) {
	testCases := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "null",
			input:   []byte("null"),
			wantErr: false,
		},
		{
			name:    "valid",
			input:   []byte(`"10MB"`),
			wantErr: false,
		},
		{
			name:    "invalid",
			input:   []byte(`"10MBB"`),
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

			q := &BandwidthQuantity{}
			err := q.UnmarshalJSON(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
