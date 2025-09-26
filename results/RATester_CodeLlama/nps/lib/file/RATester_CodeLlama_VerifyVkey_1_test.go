package file

import (
	"fmt"
	"testing"
)

func TestVerifyVkey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
