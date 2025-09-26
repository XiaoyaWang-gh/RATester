package transform

import (
	"fmt"
	"html/template"
	"testing"
)

func TestPlainify_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    template.HTML
		wantErr bool
	}{
		{
			name:  "valid input",
			input: "<p>Hello, World!</p>",
			want:  "<p>Hello, World!</p>",
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

			got, err := ns.Plainify(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Plainify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Plainify() = %v, want %v", got, tt.want)
			}
		})
	}
}
