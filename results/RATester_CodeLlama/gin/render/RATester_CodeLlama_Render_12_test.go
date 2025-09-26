package render

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// given
	html := HTML{
		Template: template.Must(template.New("").Parse("")),
		Name:     "test",
		Data:     nil,
	}

	// when
	err := html.Render(nil)

	// then
	assert.NoError(t, err)
}
