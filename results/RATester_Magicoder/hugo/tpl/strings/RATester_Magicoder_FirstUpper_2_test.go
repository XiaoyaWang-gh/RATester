package strings

import (
	"fmt"
	"testing"
)

func TestFirstUpper_2(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:  "Test with string",
			input: "hello",
			want:  "Hello",
		},
		{
			name:    "Test with non-string",
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

			got, err := ns.FirstUpper(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstUpper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FirstUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
