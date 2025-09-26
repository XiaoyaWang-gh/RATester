package models

import (
	"fmt"
	"testing"
)

func TestSet_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	j := new(JSONField)
	testData := "test_data"
	j.Set(testData)
	if *j != JSONField(testData) {
		t.Errorf("Expected %s, got %s", testData, *j)
	}
}
