package hugolib

import (
	"fmt"
	"testing"
)

func TestHugo_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	s.Hugo()
}
