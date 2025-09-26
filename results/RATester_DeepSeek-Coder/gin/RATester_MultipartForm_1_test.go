package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMultipartForm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()
	engine.MaxMultipartMemory = 1 << 20 // 1MB

	req, err := http.NewRequest("POST", "/upload", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := &Context{
		Request: req,
		engine:  engine,
	}

	_, err = ctx.MultipartForm()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
