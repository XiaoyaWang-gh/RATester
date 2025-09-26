package proxy

import (
	"fmt"
	"testing"
)

func TestClose_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &httpServer{}
	s.Close()
}
