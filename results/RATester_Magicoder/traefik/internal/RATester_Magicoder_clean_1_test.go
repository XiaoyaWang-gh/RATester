package main

import (
	"fmt"
	"testing"
)

func TestClean_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type Service struct {
		Name string `file:"-" kv:"-" label:"-"`
		Port int    `file:"-" kv:"-" label:"-"`
	}

	type TestStruct struct {
		Services map[string]*Service
	}

	testStruct := TestStruct{
		Services: map[string]*Service{
			"Service1": {
				Name: "Service1",
				Port: 8080,
			},
			"Service2": {
				Name: "Service2",
				Port: 8081,
			},
		},
	}

	clean(testStruct)

	if len(testStruct.Services) != 2 {
		t.Errorf("Expected 2 services, got %d", len(testStruct.Services))
	}

	if testStruct.Services["Service1"].Name != "Service1" || testStruct.Services["Service1"].Port != 8080 {
		t.Errorf("Expected Service1 to be cleaned, got %v", testStruct.Services["Service1"])
	}

	if testStruct.Services["Service2"].Name != "Service2" || testStruct.Services["Service2"].Port != 8081 {
		t.Errorf("Expected Service2 to be cleaned, got %v", testStruct.Services["Service2"])
	}
}
