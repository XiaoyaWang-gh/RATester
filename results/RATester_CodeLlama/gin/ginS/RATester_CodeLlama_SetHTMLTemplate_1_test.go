package ginS

import (
	"fmt"
	"html/template"
	"testing"
)

func TestSetHTMLTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// Arrange
	var templ *template.Template
	// Act
	SetHTMLTemplate(templ)
	// Assert
	// ...
}
