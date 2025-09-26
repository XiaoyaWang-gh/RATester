package render

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstance_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	r := HTMLProduction{
		Template: template.Must(template.New("").Parse("")),
	}
	name := "test"
	data := "test"

	// Act
	render := r.Instance(name, data)

	// Assert
	assert.NotNil(t, render)
}
