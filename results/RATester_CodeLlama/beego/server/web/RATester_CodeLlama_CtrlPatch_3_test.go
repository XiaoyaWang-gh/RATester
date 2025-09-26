package web

import (
	"fmt"
	"testing"
)

func TestCtrlPatch_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	rootpath := "rootpath"
	f := "f"

	// Act
	CtrlPatch(rootpath, f)

	// Assert
	// ...
}
