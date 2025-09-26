package log

import (
	"fmt"
	"testing"
)

func TestWarnf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var (
		format = "test %s"
		v      = []interface{}{"warnf"}
	)

	// Act
	Warnf(format, v...)

	// Assert
	// TODO
}
