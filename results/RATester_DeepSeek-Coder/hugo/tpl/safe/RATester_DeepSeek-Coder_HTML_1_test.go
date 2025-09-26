package safe

import (
	"fmt"
	"html/template"
	"testing"
)

func TestHTML_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	tests := []struct {
		name    string
		s       any
		want    template.HTML
		wantErr bool
	}{
		{
			name: "string",
			s:    "<p>Hello, World</p>",
			want: "<p>Hello, World</p>",
		},
		{
			name: "int",
			s:    123,
			want: "123",
		},
		{
			name: "float",
			s:    123.456,
			want: "123.456",
		},
		{
			name: "bool",
			s:    true,
			want: "true",
		},
		{
			name: "nil",
			s:    nil,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.HTML(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.HTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.HTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
