package hugofs

import (
	"fmt"
	"testing"
)

func TestOnCreate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	fs := &stacktracerFs{}
	filename := "test.txt"
	// Act
	fs.onCreate(filename)
	// Assert
	// TODO: Add assertions
}
