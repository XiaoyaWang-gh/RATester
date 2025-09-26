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

	// Arrange
	logMsgs := "test"

	// Act
	result := GetLogMsg()

	// Assert
	if result != logMsgs {
		t.Errorf("GetLogMsg() = %v, want %v", result, logMsgs)
	}
}
