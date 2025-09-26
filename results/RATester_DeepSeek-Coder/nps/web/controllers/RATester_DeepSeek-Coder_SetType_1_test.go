package controllers

import (
	"fmt"
	"testing"
)

func TestSetType_1(t *testing.T) {
	type testCase struct {
		name     string
		expected string
	}

	testCases := []testCase{
		{"test1", "test1"},
		{"test2", "test2"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			bc := &BaseController{}
			bc.SetType(tc.name)
			if bc.Data["type"] != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, bc.Data["type"])
			}
		})
	}
}
