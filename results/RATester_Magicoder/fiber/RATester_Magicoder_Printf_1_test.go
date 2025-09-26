package fiber

import (
	"fmt"
	"testing"
)

func TestPrintf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dl := &disableLogger{}
	dl.Printf("test")
}
