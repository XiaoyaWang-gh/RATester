package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetSplicedStrList_1(t *testing.T) {
	tests := []struct {
		name        string
		headerValue string
		dst         []string
		expectedDst []string
		expectedNil bool
	}{
		{
			name:        "Test Case 1",
			headerValue: "test1,test2,test3",
			dst:         make([]string, 0),
			expectedDst: []string{"test1", "test2", "test3"},
			expectedNil: false,
		},
		{
			name:        "Test Case 2",
			headerValue: "",
			dst:         make([]string, 0),
			expectedDst: nil,
			expectedNil: true,
		},
		{
			name:        "Test Case 3",
			headerValue: "test1, test2,test3",
			dst:         make([]string, 0),
			expectedDst: []string{"test1", "test2", "test3"},
			expectedNil: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			dst := getSplicedStrList(test.headerValue, test.dst)
			if test.expectedNil && dst != nil {
				t.Errorf("Expected nil, but got %v", dst)
			}
			if !test.expectedNil && !reflect.DeepEqual(dst, test.expectedDst) {
				t.Errorf("Expected %v, but got %v", test.expectedDst, dst)
			}
		})
	}
}
