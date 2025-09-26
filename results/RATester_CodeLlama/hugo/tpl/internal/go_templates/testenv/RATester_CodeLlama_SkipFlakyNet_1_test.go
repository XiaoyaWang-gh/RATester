package testenv

import (
	"fmt"
	"testing"
)

func TestSkipFlakyNet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	SkipFlakyNet(t)
}
