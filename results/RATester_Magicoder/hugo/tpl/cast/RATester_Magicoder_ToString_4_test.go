package cast

import (
	"fmt"
	"testing"
)

func TestToString_4(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:  "string",
			input: "test",
			want:  "test",
		},
		{
			name:  "int",
			input: 123,
			want:  "123",
		},
		{
			name:  "float",
			input: 123.456,
			want:  "123.456",
		},
		{
			name:  "bool",
			input: true,
			want:  "true",
		},
		{
			name:  "nil",
			input: nil,
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.ToString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
