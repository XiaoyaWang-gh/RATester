package livereload

import (
	"fmt"
	"testing"
)

func TestInitialize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	Initialize()

	// Add assertions here to verify the correctness of the method
}
