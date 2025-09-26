package alils

import (
	"fmt"
	"testing"
)

func TestSignature_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: add your test case here
}
