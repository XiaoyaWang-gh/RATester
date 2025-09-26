package gin

import (
	"fmt"
	"testing"
)

func TestSSEvent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	c := &Context{
		writermem: responseWriter{},
	}
	c.SSEvent("name", "message")
}
