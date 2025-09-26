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

	testModel := &TestStruct{
		Field1: "test",
		Field2: 123,
	}

	mi := &ModelInfo{
		FullName: "github.com/test/TestStruct",
		Model:    testModel,
	}

	mc.cache["github.com/test/TestStruct"] = mi
	mc.cacheByFullName["github.com/test/TestStruct"] = mi

	res, ok := mc.GetByMd(testModel)
	if !ok {
		t.Errorf("Expected to find model in cache")
	}

	if res != mi {
		t.Errorf("Expected %v, got %v", mi, res)
	}
}
