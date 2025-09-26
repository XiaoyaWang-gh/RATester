package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := Length{N: 5}
	expected := fmt.Sprintf(MessageTmpls["Length"], l.N)
	actual := l.DefaultMessage()
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
