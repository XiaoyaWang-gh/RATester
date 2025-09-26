package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestClean_1(t *testing.T) {
	type Service struct {
		Name string `file:"-" kv:"-" label:"-"`
		Port int    `file:"-" kv:"-" label:"-"`
	}

	type testStruct struct {
		Services map[string]*Service
	}

	testCases := []struct {
		name     string
		input    any
		expected any
	}{
		{
			name: "Test with a map of services",
			input: &testStruct{
				Services: map[string]*Service{
					"Service1": {Name: "Service1", Port: 8080},
					"Service2": {Name: "Service2", Port: 8081},
				},
			},
			expected: &testStruct{
				Services: map[string]*Service{
					"Service10": {Name: "Service1", Port: 8080},
					"Service11": {Name: "Service2", Port: 8081},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			clean(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.input)
			}
		})
	}
}
