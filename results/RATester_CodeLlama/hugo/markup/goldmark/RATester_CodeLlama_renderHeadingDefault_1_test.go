package goldmark

import (
	"fmt"
	"testing"
)

func TestRenderHeadingDefault_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
