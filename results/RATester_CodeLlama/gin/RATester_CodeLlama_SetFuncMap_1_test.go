package gin

import (
	"fmt"
	"html/template"
	"testing"
)

func TestSetFuncMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()
	engine.SetFuncMap(template.FuncMap{
		"test": func() string {
			return "test"
		},
	})
	if engine.FuncMap["test"] == nil {
		t.Error("SetFuncMap failed")
	}
}
