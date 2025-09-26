package web

import (
	"fmt"
	"testing"
)

func TestRegisterHttpServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup your test
	// TODO: test your code
	// TODO: assert results
}
