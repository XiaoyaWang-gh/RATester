package session

import (
	"fmt"
	"testing"
)

func TestSessionReleaseIfPresent_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
