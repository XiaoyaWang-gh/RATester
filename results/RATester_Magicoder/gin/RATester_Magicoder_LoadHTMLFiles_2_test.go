package gin

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin/render"
)

func TestLoadHTMLFiles_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{}

	files := []string{"file1.html", "file2.html"}
	engine.LoadHTMLFiles(files...)

	if engine.HTMLRender == nil {
		t.Error("HTMLRender should not be nil")
	}

	htmlDebug, ok := engine.HTMLRender.(render.HTMLDebug)
	if !ok {
		t.Error("HTMLRender should be of type render.HTMLDebug")
	}

	if !reflect.DeepEqual(htmlDebug.Files, files) {
		t.Error("Files in HTMLDebug should match the input files")
	}

	if !reflect.DeepEqual(htmlDebug.FuncMap, engine.FuncMap) {
		t.Error("FuncMap in HTMLDebug should match the engine's FuncMap")
	}

	if !reflect.DeepEqual(htmlDebug.Delims, engine.delims) {
		t.Error("Delims in HTMLDebug should match the engine's delims")
	}
}
