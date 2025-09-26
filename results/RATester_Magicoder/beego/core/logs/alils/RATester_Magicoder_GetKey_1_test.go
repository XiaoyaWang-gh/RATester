package alils

import (
	"fmt"
	"testing"
)

func TestGetKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "testKey"
	logContent := &LogContent{
		Key: &key,
	}

	result := logContent.GetKey()

	if result != key {
		t.Errorf("Expected %s, got %s", key, result)
	}
}
