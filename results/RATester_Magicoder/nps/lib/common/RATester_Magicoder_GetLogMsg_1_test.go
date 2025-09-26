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

	expected := "logMsgs"
	actual := GetLogMsg()
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
