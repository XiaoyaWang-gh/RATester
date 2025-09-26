package js

import (
	"fmt"
	"testing"
)

func TestTransform_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO:
}
