package postgres

import (
	"fmt"
	"testing"
)

func TestSessionRelease_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
