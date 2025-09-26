package inflect

import (
	"fmt"
	"testing"
)

func TestPluralize_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:  "single word",
			input: "apple",
			want:  "apples",
		},
		{
			name:  "multiple words",
			input: "apple tree",
			want:  "apple trees",
		},
		{
			name:    "non-string input",
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

			got, err := ns.Pluralize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pluralize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pluralize() = %v, want %v", got, tt.want)
			}
		})
	}
}
