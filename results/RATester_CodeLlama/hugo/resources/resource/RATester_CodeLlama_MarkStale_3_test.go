package resource

import (
	"fmt"
	"testing"
)

func TestMarkStale_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var os []any
	MarkStale(os...)
}
