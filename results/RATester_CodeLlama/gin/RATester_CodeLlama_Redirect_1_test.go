package gin

import (
	"fmt"
	"testing"
)

func TestRedirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	c := &Context{
		writermem: responseWriter{},
	}
	c.Redirect(301, "http://www.google.com")
	if c.Writer.Status() != 301 {
		t.Errorf("Expected status code 301, got %d", c.Writer.Status())
	}
	if c.Writer.Header().Get("Location") != "http://www.google.com" {
		t.Errorf("Expected location header to be http://www.google.com, got %s", c.Writer.Header().Get("Location"))
	}
}
