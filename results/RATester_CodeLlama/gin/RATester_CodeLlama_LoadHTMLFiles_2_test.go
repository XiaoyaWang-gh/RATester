package gin

import (
	"fmt"
	"testing"
)

func TestLoadHTMLFiles_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()
	engine.LoadHTMLFiles("./testdata/template/template1.html", "./testdata/template/template2.html")
	if engine.HTMLRender == nil {
		t.Errorf("LoadHTMLFiles should set engine.HTMLRender")
	}
}
