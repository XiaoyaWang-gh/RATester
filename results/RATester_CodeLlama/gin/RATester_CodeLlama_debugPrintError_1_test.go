package gin

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDebugPrintError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var err error
	var isDebugging bool
	var expected string
	var actual string

	// Act
	debugPrintError(err)

	// Assert
	if isDebugging {
		expected = "[GIN-debug] [ERROR] " + err.Error() + "\n"
		actual = string(DefaultErrorWriter.(*bytes.Buffer).Bytes())
		if expected != actual {
			t.Errorf("Expected %v, actual %v", expected, actual)
		}
	}
}
