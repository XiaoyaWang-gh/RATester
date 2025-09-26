package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNudge_1(t *testing.T) {
	type test struct {
		name string
		in   context
		want context
	}

	tests := []test{
		{
			name: "Test case 1",
			in:   context{state: stateTag},
			want: context{state: stateAttrName},
		},
		{
			name: "Test case 2",
			in:   context{state: stateBeforeValue},
			want: context{state: stateAttrName, delim: delimSpaceOrTagEnd, attr: attrNone},
		},
		{
			name: "Test case 3",
			in:   context{state: stateAfterName},
			want: context{state: stateAttrName, attr: attrNone},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := nudge(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
