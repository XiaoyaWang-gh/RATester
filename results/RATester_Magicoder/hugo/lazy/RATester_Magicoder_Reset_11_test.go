package lazy

import (
	"fmt"
	"testing"
)

func TestReset_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ini := &Init{}
	ini.Reset()

	// Assert that the Reset method has been called correctly
	// Add your assertions here
}
