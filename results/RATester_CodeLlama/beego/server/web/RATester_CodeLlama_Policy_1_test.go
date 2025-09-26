package web

import (
	"fmt"
	"testing"
)

func TestPolicy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pattern := "pattern"
	method := "method"
	policy := []PolicyFunc{}
	Policy(pattern, method, policy...)
}
