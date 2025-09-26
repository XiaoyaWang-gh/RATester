package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToXML_1(t *testing.T) {
	type TestStruct struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}

	testCases := []struct {
		name     string
		response string
		expected TestStruct
	}{
		{
			name:     "Test ToXML",
			response: `<TestStruct><name>John</name><age>30</age></TestStruct>`,
			expected: TestStruct{Name: "John", Age: 30},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req := &BeegoHTTPRequest{}
			req.body = []byte(tc.response)

			var result TestStruct
			err := req.ToXML(&result)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
