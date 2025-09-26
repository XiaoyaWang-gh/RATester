package orm

import (
	"fmt"
	"testing"
)

func TestAndCond_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Condition{}
	cond := &Condition{}
	c.AndCond(cond)
}
