package urlreplacers

import (
	"fmt"
	"testing"
)

func TestNewPrefixState_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
