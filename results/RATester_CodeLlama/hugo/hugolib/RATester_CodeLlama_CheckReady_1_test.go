package hugolib

import (
	"fmt"
	"testing"
)

func TestCheckReady_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	s.state = siteStateReady
	s.CheckReady()
}
