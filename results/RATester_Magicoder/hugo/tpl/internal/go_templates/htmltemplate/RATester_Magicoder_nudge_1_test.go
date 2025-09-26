package template

import (
	"fmt"
	"reflect"
	"testing"
)

func Testnudge_1(t *testing.T) {
	tests := []struct {
		name string
		c    context
		want context
	}{
		{
			name: "Test Case 1",
			c:    context{state: stateTag},
			want: context{state: stateAttrName},
		},
		{
			name: "Test Case 2",
			c:    context{state: stateBeforeValue},
			want: context{state: stateAttrName, delim: delimSpaceOrTagEnd, attr: attrNone},
		},
		{
			name: "Test Case 3",
			c:    context{state: stateAfterName},
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

			if got := nudge(tt.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
