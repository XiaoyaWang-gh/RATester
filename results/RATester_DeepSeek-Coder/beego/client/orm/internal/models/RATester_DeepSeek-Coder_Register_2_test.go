package models

import (
	"fmt"
	"testing"
)

func TestRegister_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{
		cache:           make(map[string]*ModelInfo),
		cacheByFullName: make(map[string]*ModelInfo),
	}

	type TestModel struct {
		Id   int
		Name string
	}

	err := mc.Register("", false, &TestModel{})
	if err != nil {
		t.Errorf("Register error: %v", err)
	}

	mi, ok := mc.Get("TestModel")
	if !ok {
		t.Errorf("Get error: model not found")
	}

	if mi.Table != "TestModel" {
		t.Errorf("Register error: table name mismatch, expected 'TestModel', got '%s'", mi.Table)
	}

	if mi.Pkg != "main" {
		t.Errorf("Register error: package name mismatch, expected 'main', got '%s'", mi.Pkg)
	}

	if mi.Model.(*TestModel) == nil {
		t.Errorf("Register error: model mismatch, expected 'TestModel', got nil")
	}
}
