package hugolib

import (
	"fmt"
	"testing"
)

func TestResources_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageState{}
	p.s = &Site{}
	p.s.pageMap = &pageMap{}
	p.s.pageMap.getOrCreateResourcesForPage(p)
}
