package crypto

import (
	"fmt"
	"testing"
)

func TestFNV32a_2(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    int
		wantErr bool
	}{
		{
			name:  "valid input",
			input: "test",
			want:  305419896,
		},
		{
			name:    "invalid input",
			input:   123,
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
