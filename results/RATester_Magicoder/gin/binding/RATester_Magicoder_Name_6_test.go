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
	name := b.Name()
	if name != "plain" {
		t.Errorf("Expected 'plain', but got '%s'", name)
	}
}
