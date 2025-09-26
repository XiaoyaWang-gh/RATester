package gin

import (
	"fmt"
	"html/template"
	"testing"
)

func TestSetHTMLTemplate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()
	engine.SetHTMLTemplate(template.Must(template.New("").Parse("")))
	if engine.HTMLRender == nil {
		t.Errorf("SetHTMLTemplate() did not set HTMLRender")
	}
}
