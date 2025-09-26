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
	expected := "msgpack"
	actual := b.Name()
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
