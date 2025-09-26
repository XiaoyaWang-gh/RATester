package modules

import (
	"fmt"
	"testing"
)

func TestCollect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &collector{}
	c.collect()
}
