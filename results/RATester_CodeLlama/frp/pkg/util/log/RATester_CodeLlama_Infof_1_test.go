package log

import (
	"fmt"
	"testing"
)

func TestInfof_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var (
		format = "test %s"
		v      = []interface{}{"test"}
	)

	// Act
	Infof(format, v...)

	// Assert
	// TODO
}
