package csrf

import (
	"fmt"
	"testing"
)

func TestUnmarshalMsg_5(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "valid input",
			input:   []byte(`{"field": "value"}`),
			wantErr: false,
		},
		{
			name:    "invalid input",
			input:   []byte(`{"field": invalid}`),
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []byte{},
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

			var m storageManager
			_, err := m.UnmarshalMsg(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("UnmarshalMsg() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
