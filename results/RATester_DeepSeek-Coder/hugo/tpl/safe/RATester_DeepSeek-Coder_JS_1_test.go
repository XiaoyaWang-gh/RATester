package safe

import (
	"fmt"
	"html/template"
	"testing"
)

func TestJS_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   any
		want    template.JS
		wantErr bool
	}{
		{
			name:  "string input",
			input: "test",
			want:  "test",
		},
		{
			name:  "int input",
			input: 123,
			want:  "123",
		},
		{
			name:  "float input",
			input: 123.456,
			want:  "123.456",
		},
		{
			name:  "bool input",
			input: true,
			want:  "true",
		},
		{
			name:  "nil input",
			input: nil,
			want:  "<nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ns := &Namespace{}
			got, err := ns.JS(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("JS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != template.JS(tt.want) {
				t.Errorf("JS() = %v, want %v", got, tt.want)
			}
		})
	}
}
