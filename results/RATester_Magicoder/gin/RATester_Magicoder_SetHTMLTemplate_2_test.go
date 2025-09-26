package gin

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/gin-gonic/gin/render"
)

func TestSetHTMLTemplate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new Engine instance
	engine := New()

	// Create a new template
	templ := template.New("test")

	// Set the template to the engine
	engine.SetHTMLTemplate(templ)

	// Check if the HTMLRender is set correctly
	if engine.HTMLRender == nil {
		t.Error("HTMLRender is not set correctly")
	}

	// Check if the template is set correctly
	if engine.HTMLRender.(render.HTMLProduction).Template != templ.Funcs(engine.FuncMap) {
		t.Error("Template is not set correctly")
	}
}
