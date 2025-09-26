package math

import (
	"fmt"
	"testing"
)

func TestMod_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		n1      any
		n2      any
		want    int64
		wantErr bool
	}{
		{
			name:    "Test case 1",
			n1:      10,
			n2:      2,
			want:    0,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			n1:      10.5,
			n2:      2,
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test case 3",
			n1:      10,
			n2:      0,
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test case 4",
			n1:      10,
			n2:      "2",
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

			got, err := ns.Mod(tt.n1, tt.n2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Mod() = %v, want %v", got, tt.want)
			}
		})
	}
}
