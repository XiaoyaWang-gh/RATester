package math

import (
	"fmt"
	"testing"
)

func TestLog_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    float64
		wantErr bool
	}{
		{
			name:    "Test with valid input",
			input:   10.0,
			want:    2.302585092994046,
			wantErr: false,
		},
		{
			name:    "Test with invalid input",
			input:   "abc",
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

			got, err := ns.Log(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Log() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}
