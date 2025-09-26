package cast

import (
	"fmt"
	"testing"
)

func TestToInt_2(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    int
		wantErr bool
	}{
		{
			name:  "int",
			input: 123,
			want:  123,
		},
		{
			name:  "float64",
			input: 123.456,
			want:  123,
		},
		{
			name:  "string",
			input: "123",
			want:  123,
		},
		{
			name:    "invalid string",
			input:   "abc",
			wantErr: true,
		},
		{
			name:    "nil",
			input:   nil,
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

			got, err := ns.ToInt(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
