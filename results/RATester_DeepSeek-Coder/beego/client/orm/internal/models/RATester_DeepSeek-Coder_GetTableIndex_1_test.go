package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetTableIndex_1(t *testing.T) {
	type testStruct struct {
		TableIndex func() [][]string
	}

	testCases := []struct {
		name     string
		input    testStruct
		expected [][]string
	}{
		{
			name: "Test case 1",
			input: testStruct{
				TableIndex: func() [][]string {
					return [][]string{{"1", "2"}, {"3", "4"}}
				},
			},
			expected: [][]string{{"1", "2"}, {"3", "4"}},
		},
		{
			name: "Test case 2",
			input: testStruct{
				TableIndex: func() [][]string {
					return [][]string{{"5", "6"}, {"7", "8"}}
				},
			},
			expected: [][]string{{"5", "6"}, {"7", "8"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			val := reflect.ValueOf(tc.input)
			result := GetTableIndex(val)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
