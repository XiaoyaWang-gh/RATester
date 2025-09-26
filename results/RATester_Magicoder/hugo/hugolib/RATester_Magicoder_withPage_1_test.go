package hugolib

import (
	"fmt"
	"testing"
)

func TestwithPage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoSites{}
	fn := func(s string, p *pageState) bool {
		// Your test case logic here
		return true
	}
	h.withPage(fn)
}
