package log

import (
	"fmt"
	"testing"
)

func TestPrint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	info := &logInfo{}
	log := func(f interface{}, v ...interface{}) {}
	// Act
	info.print(log)
	// Assert
}
