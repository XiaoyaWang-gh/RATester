package cast

import (
	"fmt"
	"testing"
)

func TestToFloat_2(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    float64
		wantErr bool
	}{
		{
			name:    "Test case 1",
			input:   "123.45",
			want:    123.45,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test case 3",
			input:   nil,
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

			got, err := ns.ToFloat(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
