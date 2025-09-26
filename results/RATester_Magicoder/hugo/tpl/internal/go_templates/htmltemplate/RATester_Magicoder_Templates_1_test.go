package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTemplates_1(t *testing.T) {
	ns := &nameSpace{
		set: make(map[string]*Template),
	}

	t1 := &Template{
		nameSpace: ns,
	}
	t2 := &Template{
		nameSpace: ns,
	}

	ns.set["t1"] = t1
	ns.set["t2"] = t2

	tests := []struct {
		name string
		want []*Template
	}{
		{
			name: "Test case 1",
			want: []*Template{t1, t2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := t1.Templates()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Templates() = %v, want %v", got, tt.want)
			}
		})
	}
}
