package template

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestLookup_1(t *testing.T) {
	type testStruct struct {
		name string
		want *Template
	}

	tests := []testStruct{
		{
			name: "Test case 1",
			want: &Template{
				nameSpace: &nameSpace{
					mu:      sync.Mutex{},
					set:     map[string]*Template{},
					escaped: false,
				},
			},
		},
		{
			name: "Test case 2",
			want: &Template{
				nameSpace: &nameSpace{
					mu:      sync.Mutex{},
					set:     map[string]*Template{"test": &Template{}},
					escaped: false,
				},
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
				nameSpace: &nameSpace{
					mu:      sync.Mutex{},
					set:     map[string]*Template{},
					escaped: false,
				},
			}

			got := temp.Lookup(tt.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
