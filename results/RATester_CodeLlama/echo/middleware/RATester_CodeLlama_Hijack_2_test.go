package middleware

import (
	"fmt"
	"testing"
)

func TestHijack_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
