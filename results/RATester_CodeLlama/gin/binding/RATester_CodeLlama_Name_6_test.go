package binding

import (
	"fmt"
	"testing"
)

func TestName_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := plainBinding{}
	if b.Name() != "plain" {
		t.Errorf("Expected plain, got %s", b.Name())
	}
}
