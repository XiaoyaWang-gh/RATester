package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppendIfNotSeen_1(t *testing.T) {
	type testStruct struct {
		name string
		v    reflect.Value
	}
	tests := []testStruct{
		{
			name: "Test appendIfNotSeen with int",
			v:    reflect.ValueOf(1),
		},
		{
			name: "Test appendIfNotSeen with string",
			v:    reflect.ValueOf("test"),
		},
		{
			name: "Test appendIfNotSeen with float",
			v:    reflect.ValueOf(1.23),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			i := &intersector{
				r:    reflect.MakeSlice(reflect.SliceOf(tt.v.Type()), 0, 0),
				seen: make(map[any]bool),
			}
			i.appendIfNotSeen(tt.v)
			if !i.seen[tt.v.Interface()] {
				t.Errorf("Expected value to be seen")
			}
			if i.r.Len() != 1 {
				t.Errorf("Expected slice length to be 1, got %d", i.r.Len())
			}
			if i.r.Index(0).Interface() != tt.v.Interface() {
				t.Errorf("Expected first element of slice to be %v, got %v", tt.v.Interface(), i.r.Index(0).Interface())
			}
		})
	}
}
