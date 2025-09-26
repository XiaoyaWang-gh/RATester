package common

import (
	"fmt"
	"testing"
)

func TestGetLogMsg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := "log messages"
	actual := GetLogMsg()
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
