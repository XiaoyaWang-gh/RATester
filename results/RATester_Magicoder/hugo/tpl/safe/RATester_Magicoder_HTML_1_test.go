package safe

import (
	"fmt"
	"html/template"
	"testing"
)

func TestHTML_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    template.HTML
		wantErr bool
	}{
		{
			name:  "string",
			input: "<script>alert('hello')</script>",
			want:  "&lt;script&gt;alert('hello')&lt;/script&gt;",
		},
		{
			name:  "int",
			input: 123,
			want:  "123",
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

			got, err := ns.HTML(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
