package gin

import (
	"fmt"
	"testing"
)

func TestLoadHTMLGlob_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	engine := New()
	engine.LoadHTMLGlob("testdata/glob/*")
	if engine.HTMLRender == nil {
		t.Error("HTMLRender should not be nil")
	}
}
