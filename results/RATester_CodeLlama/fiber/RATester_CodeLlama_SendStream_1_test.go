package fiber

import (
	"fmt"
	"testing"
)

func TestSendStream_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
