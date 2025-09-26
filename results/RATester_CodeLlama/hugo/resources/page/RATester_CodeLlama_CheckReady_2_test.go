package page

import (
	"fmt"
	"testing"
)

func TestCheckReady_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := testSite{}
	s.CheckReady()
}
