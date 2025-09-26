package math

import (
	"fmt"
	"testing"
)

func TestToDegrees_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    float64
		wantErr bool
	}{
		{
			name:    "Test with valid input",
			input:   1,
			want:    57.29577951308232,
			wantErr: false,
		},
		{
			name:    "Test with invalid input",
			input:   "invalid",
			want:    0,
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

			got, err := ns.ToDegrees(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDegrees() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToDegrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
