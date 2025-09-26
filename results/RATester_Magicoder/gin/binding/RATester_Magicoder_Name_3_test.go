package binding

import (
	"fmt"
	"testing"
)

func TestName_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := jsonBinding{}
	name := b.Name()
	if name != "json" {
		t.Errorf("Expected 'json', but got '%s'", name)
	}
}
