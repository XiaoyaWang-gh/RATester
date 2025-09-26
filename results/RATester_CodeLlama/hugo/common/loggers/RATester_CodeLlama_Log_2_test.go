package loggers

import (
	"fmt"
	"testing"
)

func TestLog_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var log Logger
	// Act
	log = Log()
	// Assert
	if log == nil {
		t.Errorf("Log() failed")
	}
}
