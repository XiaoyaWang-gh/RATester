package gin

import (
	"fmt"
	"html/template"
	"testing"
)

func TestDebugPrintLoadTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	tmpl := &template.Template{}
	// Act
	debugPrintLoadTemplate(tmpl)
	// Assert
	// TODO: Add assertions
}
