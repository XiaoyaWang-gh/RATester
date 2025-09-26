package gin

import (
	"fmt"
	"testing"
)

func TestRunFd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: add test cases
}
