package nathole

import (
	"fmt"
	"testing"
)

func TestGenSid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	sid := c.GenSid()
	t.Log(sid)
}
