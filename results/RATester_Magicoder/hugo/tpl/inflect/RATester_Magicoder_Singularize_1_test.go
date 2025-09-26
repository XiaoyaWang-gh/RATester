package inflect

import (
	"fmt"
	"testing"
)

func TestSingularize_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:    "Test singularize with a string",
			input:   "cats",
			want:    "cat",
			wantErr: false,
		},
		{
			name:    "Test singularize with a number",
			input:   123,
			want:    "",
			wantErr: true,
		},
		{
			name:    "Test singularize with a nil",
			input:   nil,
			want:    "",
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

			got, err := ns.Singularize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Singularize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Singularize() = %v, want %v", got, tt.want)
			}
		})
	}
}
