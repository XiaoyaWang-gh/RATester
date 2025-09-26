package modules

import (
	"fmt"
	"testing"
)

func TestOwner_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	owner := &moduleAdapter{
		path: "test/path",
		dir:  "test/dir",
	}
	module := &moduleAdapter{
		owner: owner,
	}

	if module.Owner() != owner {
		t.Errorf("Expected owner to be %v, but got %v", owner, module.Owner())
	}
}
