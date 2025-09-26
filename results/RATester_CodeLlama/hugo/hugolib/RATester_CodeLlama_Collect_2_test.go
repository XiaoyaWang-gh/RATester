package hugolib

import (
	"fmt"
	"testing"
)

func TestCollect_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &pagesCollector{}
	c.Collect()
}
