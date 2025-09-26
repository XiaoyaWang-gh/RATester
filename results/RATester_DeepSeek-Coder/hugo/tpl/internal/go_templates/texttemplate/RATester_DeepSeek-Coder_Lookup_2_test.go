package template

import (
	"fmt"
	"testing"
)

func TestLookup_2(t *testing.T) {
	type testStruct struct {
		name string
		want *Template
	}

	tests := []testStruct{
		{
			name: "Test case 1",
			want: &Template{
				name: "test",
			},
		},
		{
			name: "Test case 2",
			want: &Template{
				name: "test2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			temp := &Template{
				name: "test",
			}
			temp.Lookup(tt.name)
			if temp.Lookup(tt.name) != tt.want {
				t.Errorf("Lookup() = %v, want %v", temp.Lookup(tt.name), tt.want)
			}
		})
	}
}
