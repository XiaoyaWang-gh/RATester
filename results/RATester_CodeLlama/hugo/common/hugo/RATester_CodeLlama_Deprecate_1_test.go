package hugo

import (
	"fmt"
	"testing"
)

func TestDeprecate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	item := "item"
	alternative := "alternative"
	version := "1.0.0"
	Deprecate(item, alternative, version)
}
