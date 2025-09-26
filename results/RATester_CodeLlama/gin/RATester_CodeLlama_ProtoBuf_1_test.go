package gin

import (
	"fmt"
	"testing"
)

func TestProtoBuf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	var c Context
	c.ProtoBuf(200, "test")
}
