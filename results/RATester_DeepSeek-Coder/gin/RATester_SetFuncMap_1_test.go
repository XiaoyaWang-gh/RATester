package gin

import (
	"fmt"
	"html/template"
	"strings"
	"testing"
)

func TestSetFuncMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new engine
	engine := New()

	// Define a function map
	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
	}

	// Set the function map
	engine.SetFuncMap(funcMap)

	// Check if the function map was set correctly
	if engine.FuncMap["upper"] != funcMap["upper"] {
		t.Errorf("Expected upper function to be set")
	}
	if engine.FuncMap["lower"] != funcMap["lower"] {
		t.Errorf("Expected lower function to be set")
	}
}
