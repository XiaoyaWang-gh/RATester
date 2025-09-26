package context

import (
	"fmt"
	"testing"
)

func TestContentType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{}

	// Test with a valid extension
	output.ContentType(".html")
	if output.Context.ResponseWriter.Header().Get("Content-Type") != "text/html; charset=utf-8" {
		t.Error("Expected Content-Type to be 'text/html; charset=utf-8', but got", output.Context.ResponseWriter.Header().Get("Content-Type"))
	}

	// Test with an invalid extension
	output.ContentType(".invalid")
	if output.Context.ResponseWriter.Header().Get("Content-Type") != "" {
		t.Error("Expected Content-Type to be '', but got", output.Context.ResponseWriter.Header().Get("Content-Type"))
	}
}
