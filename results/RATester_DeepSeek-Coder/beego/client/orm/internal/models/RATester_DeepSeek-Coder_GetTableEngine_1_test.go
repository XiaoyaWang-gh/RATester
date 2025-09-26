package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetTableEngine_1(t *testing.T) {
	type testStruct struct {
		tableEngine string
	}

	testCases := []struct {
		name   string
		input  testStruct
		expect string
	}{
		{
			name:   "Test case 1",
			input:  testStruct{tableEngine: "InnoDB"},
			expect: "InnoDB",
		},
		{
			name:   "Test case 2",
			input:  testStruct{tableEngine: "MyISAM"},
			expect: "MyISAM",
		},
		{
			name:   "Test case 3",
			input:  testStruct{tableEngine: "Memory"},
			expect: "Memory",
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
			result := GetTableEngine(val)
			if result != tc.expect {
				t.Errorf("Expected %s, but got %s", tc.expect, result)
			}
		})
	}
}
