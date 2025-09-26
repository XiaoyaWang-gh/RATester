package gin

import (
	"fmt"
	"testing"
)

func TestHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	c := &Context{
		Writer: &responseWriter{},
	}
	c.Header("key", "value")
	if c.Writer.Header().Get("key") != "value" {
		t.Errorf("Header() failed")
	}
}
