package strings

import (
	"fmt"
	"html/template"
	"reflect"
	"testing"
)

func TestChomp_2(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		s       any
		want    any
		wantErr bool
	}{
		{
			name:    "Test with string",
			s:       "Hello, World!",
			want:    "Hello, World",
			wantErr: false,
		},
		{
			name:    "Test with HTML",
			s:       template.HTML("<p>Hello, World!</p>"),
			want:    template.HTML("<p>Hello, World</p>"),
			wantErr: false,
		},
		{
			name:    "Test with non-string",
			s:       123,
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

			got, err := ns.Chomp(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chomp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chomp() = %v, want %v", got, tt.want)
			}
		})
	}
}
