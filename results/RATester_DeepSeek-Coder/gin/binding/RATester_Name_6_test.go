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
	expected := "plain"
	actual := b.Name()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
