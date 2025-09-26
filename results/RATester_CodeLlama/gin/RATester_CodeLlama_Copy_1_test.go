package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCopy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		writermem: responseWriter{},
		Request:   &http.Request{},
		engine:    &Engine{},
	}
	cp := c.Copy()
	if cp.writermem != c.writermem {
		t.Errorf("Expected %v, got %v", c.writermem, cp.writermem)
	}
	if cp.Request != c.Request {
		t.Errorf("Expected %v, got %v", c.Request, cp.Request)
	}
	if cp.engine != c.engine {
		t.Errorf("Expected %v, got %v", c.engine, cp.engine)
	}
}
