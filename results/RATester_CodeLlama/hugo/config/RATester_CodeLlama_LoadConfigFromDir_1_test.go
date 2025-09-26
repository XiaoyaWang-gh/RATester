package config

import (
	"fmt"
	"testing"
)

func TestLoadConfigFromDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// TODO(bep)
}
