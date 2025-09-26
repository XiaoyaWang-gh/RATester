package models

import (
	"fmt"
	"testing"
)

func TestGetByMd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{
		cache:           make(map[string]*ModelInfo),
		cacheByFullName: make(map[string]*ModelInfo),
	}

	type TestStruct struct {
		Field1 string
		Field2 int
	}

	testStruct := TestStruct{
		Field1: "test",
		Field2: 123,
	}

	mi, ok := mc.GetByMd(testStruct)
	if !ok {
		t.Error("Expected model info, got nil")
	}

	if mi.FullName != "TestStruct" {
		t.Errorf("Expected full name 'TestStruct', got '%s'", mi.FullName)
	}
}
