package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToMarshalJSON_1(t *testing.T) {
	type testCase struct {
		name     string
		receiver string
		pkgPath  string
		excludes []string
		want     string
		want1    []string
	}

	tests := []testCase{
		{
			name:     "Test case 1",
			receiver: "s",
			pkgPath:  "github.com/gohugoio/hugo/codegen",
			excludes: []string{"^ExcludeThis$"},
			want:     "expected output",
			want1:    []string{"encoding/json"},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := Methods{
				{Name: "IncludeThis", In: []string{"string"}, Out: []string{"string", "error"}},
				{Name: "ExcludeThis", In: []string{"string"}, Out: []string{"string", "error"}},
			}

			got, got1 := m.ToMarshalJSON(tt.receiver, tt.pkgPath, tt.excludes...)
			if got != tt.want {
				t.Errorf("ToMarshalJSON() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToMarshalJSON() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
