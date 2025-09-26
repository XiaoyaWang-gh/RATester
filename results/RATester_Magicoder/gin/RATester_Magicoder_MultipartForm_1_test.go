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
	engine.MaxMultipartMemory = 32 << 20 // 32MB

	req, err := http.NewRequest("POST", "/upload", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := &Context{
		Request: req,
		engine:  engine,
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		t.Fatal(err)
	}

	if form == nil {
		t.Fatal("Expected form to be not nil")
	}
}
