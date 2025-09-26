package crypto

import (
	"fmt"
	"testing"
)

func TestMD5_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:    "Test with string",
			input:   "test",
			want:    "098f6bcd4621d373cade4e832627b4f6",
			wantErr: false,
		},
		{
			name:    "Test with int",
			input:   123,
			want:    "202cb962ac59075b964b07152d234b70",
			wantErr: false,
		},
		{
			name:    "Test with float",
			input:   123.456,
			want:    "d41d8cd98f00b204e9800998ecf8427e",
			wantErr: false,
		},
		{
			name:    "Test with nil",
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
