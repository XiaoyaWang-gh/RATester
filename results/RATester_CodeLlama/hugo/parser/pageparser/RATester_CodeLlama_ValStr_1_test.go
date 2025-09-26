package pageparser

import (
	"fmt"
	"testing"
)

func TestValStr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var i Item
	var source []byte

	// Act
	var actual string
	actual = i.ValStr(source)

	// Assert
	var expected string
	expected = ""
	if actual != expected {
		t.Errorf("Expected %v, but actual is %v", expected, actual)
		return
	}
}
