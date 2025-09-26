package hugolib

import (
	"fmt"
	"testing"
)

func TestString_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageState{}
	p.Path()
	p.String()
}
