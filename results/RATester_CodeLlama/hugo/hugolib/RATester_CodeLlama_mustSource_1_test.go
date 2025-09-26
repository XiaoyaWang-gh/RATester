package hugolib

import (
	"fmt"
	"testing"
)

func TestMustSource_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &cachedContent{}
	c.mustSource()
}
