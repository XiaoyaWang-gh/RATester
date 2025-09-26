package hugolib

import (
	"fmt"
	"testing"
)

func TestNextPage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := pagePosition{}
	p.NextPage()
}
