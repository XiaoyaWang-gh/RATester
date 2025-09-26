package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestcontextAfterText_1(t *testing.T) {
	tests := []struct {
		name string
		c    context
		s    []byte
		want context
	}{
		{
			name: "Test case 1",
			c:    context{state: stateTag, delim: delimNone},
			s:    []byte("test"),
			want: context{state: stateTag, delim: delimNone},
		},
		{
			name: "Test case 2",
			c:    context{state: stateTag, delim: delimSpaceOrTagEnd},
			s:    []byte("test"),
			want: context{state: stateTag, delim: delimSpaceOrTagEnd},
		},
		{
			name: "Test case 3",
			c:    context{state: stateTag, delim: delimSpaceOrTagEnd},
			s:    []byte("test"),
			want: context{state: stateTag, delim: delimSpaceOrTagEnd},
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

			if got, _ := contextAfterText(tt.c, tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("contextAfterText() = %v, want %v", got, tt.want)
			}
		})
	}
}
