package math

import (
	"fmt"
	"testing"
)

func TestModBool_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		n1      any
		n2      any
		want    bool
		wantErr bool
	}{
		{
			name:    "Test case 1",
			n1:      10,
			n2:      2,
			want:    false,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			n1:      15,
			n2:      3,
			want:    true,
			wantErr: false,
		},
		{
			name:    "Test case 3",
			n1:      20,
			n2:      0,
			want:    false,
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

			got, err := ns.ModBool(tt.n1, tt.n2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ModBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
