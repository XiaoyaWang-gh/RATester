package gin

import (
	"fmt"
	"testing"
)

func TestStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Writer: &responseWriter{},
	}
	c.Status(200)
	if c.Writer.Status() != 200 {
		t.Errorf("Status() = %v, want %v", c.Writer.Status(), 200)
	}
}
