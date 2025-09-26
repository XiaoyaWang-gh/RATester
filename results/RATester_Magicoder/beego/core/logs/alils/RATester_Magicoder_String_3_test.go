package alils

import (
	"fmt"
	"testing"
)

func TestString_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "testKey"
	value := "testValue"
	logContent := &LogContent{
		Key:   &key,
		Value: &value,
	}
	expected := fmt.Sprintf("Key: %s\nValue: %s\n", key, value)
	if logContent.String() != expected {
		t.Errorf("Expected %s, but got %s", expected, logContent.String())
	}
}
