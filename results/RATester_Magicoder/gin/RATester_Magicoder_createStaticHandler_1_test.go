package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestcreateStaticHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	group := &RouterGroup{}
	relativePath := "/static"
	fs := http.Dir("./static")
	handler := group.createStaticHandler(relativePath, fs)

	// Test case 1: File exists
	context := &Context{
		Params: Params{
			Param{Key: "filepath", Value: "test.txt"},
		},
	}
	handler(context)
	if context.Writer.Status() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, context.Writer.Status())
	}

	// Test case 2: File does not exist
	context = &Context{
		Params: Params{
			Param{Key: "filepath", Value: "nonexistent.txt"},
		},
	}
	handler(context)
	if context.Writer.Status() != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, context.Writer.Status())
	}
}
