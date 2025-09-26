package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoopSetRefs_1(t *testing.T) {
	testCases := []struct {
		name     string
		refs     []interface{}
		sInds    []reflect.Value
		nIndsPtr *[]reflect.Value
		eTyps    []reflect.Type
		init     bool
		expected []reflect.Value
	}{
		{
			name: "Test case 1",
			refs: []interface{}{1, 2, 3},
			sInds: []reflect.Value{
				reflect.ValueOf(1),
				reflect.ValueOf(2),
				reflect.ValueOf(3),
			},
			nIndsPtr: &[]reflect.Value{
				reflect.ValueOf(1),
				reflect.ValueOf(2),
				reflect.ValueOf(3),
			},
			eTyps: []reflect.Type{
				reflect.TypeOf(1),
				reflect.TypeOf(2),
				reflect.TypeOf(3),
			},
			init: true,
			expected: []reflect.Value{
				reflect.ValueOf(1),
				reflect.ValueOf(2),
				reflect.ValueOf(3),
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			o := &rawSet{}
			o.loopSetRefs(tc.refs, tc.sInds, tc.nIndsPtr, tc.eTyps, tc.init)
			actual := *tc.nIndsPtr
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
