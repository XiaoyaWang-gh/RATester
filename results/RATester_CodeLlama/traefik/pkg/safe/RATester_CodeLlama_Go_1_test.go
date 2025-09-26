package safe

import (
	"fmt"
	"testing"
)

func TestGo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var goroutine func()
	Go(goroutine)
}
