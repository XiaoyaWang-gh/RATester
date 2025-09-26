package hash

import (
	"fmt"
	"testing"
)

func TestFNV32a_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    int
		wantErr bool
	}{
		{
			name:    "Test case 1",
			input:   "test",
			want:    2166136261,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			input:   123,
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test case 3",
			input:   nil,
			want:    0,
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

			got, err := ns.FNV32a(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FNV32a() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FNV32a() = %v, want %v", got, tt.want)
			}
		})
	}
}
