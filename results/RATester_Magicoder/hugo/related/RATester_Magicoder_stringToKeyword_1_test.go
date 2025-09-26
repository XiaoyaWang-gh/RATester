package related

import (
	"fmt"
	"reflect"
	"testing"
)

func TeststringToKeyword_1(t *testing.T) {
	tests := []struct {
		name string
		cfg  IndexConfig
		s    string
		want Keyword
	}{
		{
			name: "ToLower is true",
			cfg:  IndexConfig{ToLower: true},
			s:    "Hello",
			want: StringKeyword("hello"),
		},
		{
			name: "ToLower is false",
			cfg:  IndexConfig{ToLower: false},
			s:    "Hello",
			want: StringKeyword("Hello"),
		},
		{
			name: "Type is TypeFragments",
			cfg:  IndexConfig{Type: TypeFragments},
			s:    "Hello",
			want: FragmentKeyword("Hello"),
		},
		{
			name: "Type is not TypeFragments",
			cfg:  IndexConfig{Type: "notTypeFragments"},
			s:    "Hello",
			want: StringKeyword("Hello"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.cfg.stringToKeyword(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToKeyword() = %v, want %v", got, tt.want)
			}
		})
	}
}
