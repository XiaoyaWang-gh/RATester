package es

import (
	"fmt"
	"testing"
)

func TestDestroy_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	el := &esLogger{}
	el.Destroy()
}
