package client

import (
	"fmt"
	"testing"
)

func TestClose_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{}
	svr.Close()
}
