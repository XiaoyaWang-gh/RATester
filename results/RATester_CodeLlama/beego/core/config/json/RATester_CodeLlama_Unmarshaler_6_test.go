package json

import (
	"fmt"
	"testing"
)

func TestUnmarshaler_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
