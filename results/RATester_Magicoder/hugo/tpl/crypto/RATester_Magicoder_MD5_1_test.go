package crypto

import (
	"fmt"
	"testing"
)

func TestMD5_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			input:   "test",
			want:    "098f6bcd4621d373cade4e832627b4f6",
			wantErr: false,
		},
		{
			name:    "Test Case 2",
			input:   123,
			want:    "",
			wantErr: true,
		},
		{
			name:    "Test Case 3",
			input:   nil,
			want:    "d41d8cd98f00b204e9800998ecf8427e",
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

			got, err := ns.MD5(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MD5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}
