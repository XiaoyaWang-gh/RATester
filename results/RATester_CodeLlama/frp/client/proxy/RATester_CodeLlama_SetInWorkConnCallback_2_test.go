package proxy

import (
	"fmt"
	"testing"
)

func TestSetInWorkConnCallback_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
