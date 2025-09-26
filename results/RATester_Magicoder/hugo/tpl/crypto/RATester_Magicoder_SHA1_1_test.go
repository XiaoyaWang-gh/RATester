package crypto

import (
	"fmt"
	"testing"
)

func TestSHA1_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			input:   "test",
			want:    "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			input:   123,
			want:    "",
			wantErr: true,
		},
		{
			name:    "Test case 3",
			input:   nil,
			want:    "da39a3ee5e6b4b0d3255bfef95601890afd80709",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.SHA1(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SHA1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}
