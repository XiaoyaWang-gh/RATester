package commands

import (
	"fmt"
	"testing"
)

func TestInitMemTicker_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &hugoBuilder{}
	f := c.initMemTicker()
	f()
}
