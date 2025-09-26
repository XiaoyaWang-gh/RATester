package web

import (
	"fmt"
	"testing"
)

func TestCtrlPost_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	rootpath := "rootpath"
	f := "f"
	// Act
	CtrlPost(rootpath, f)
	// Assert
	// ...
}
