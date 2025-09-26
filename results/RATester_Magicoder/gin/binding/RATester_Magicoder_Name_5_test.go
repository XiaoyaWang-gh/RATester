package binding

import (
	"fmt"
	"testing"
)

func TestName_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := msgpackBinding{}
	name := b.Name()
	if name != "msgpack" {
		t.Errorf("Expected 'msgpack', but got '%s'", name)
	}
}
