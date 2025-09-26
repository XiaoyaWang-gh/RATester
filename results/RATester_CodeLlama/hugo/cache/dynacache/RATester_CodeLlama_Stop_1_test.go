package dynacache

import (
	"fmt"
	"testing"
)

func TestStop_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache{}
	c.Stop()
}
