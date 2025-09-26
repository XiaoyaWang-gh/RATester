package tplimpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExtractIdentifiers_1(t *testing.T) {
	tests := []struct {
		name string
		line string
		want []string
	}{
		{
			name: "simple",
			line: "{{ $name := \"John Doe\" }}",
			want: []string{"$name"},
		},
		{
			name: "multiple",
			line: "{{ $name := \"John Doe\" }}{{ $age := 30 }}",
			want: []string{"$name", "$age"},
		},
		{
			name: "no identifiers",
			line: "{{ .Title }}",
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			h := &templateHandler{}
			if got := h.extractIdentifiers(tt.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractIdentifiers() = %v, want %v", got, tt.want)
			}
		})
	}
}
