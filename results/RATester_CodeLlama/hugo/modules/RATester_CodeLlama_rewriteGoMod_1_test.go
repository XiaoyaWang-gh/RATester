package modules

import (
	"fmt"
	"testing"
)

func TestRewriteGoMod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO(bep)
}
