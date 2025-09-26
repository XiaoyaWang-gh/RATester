package binding

import (
	"fmt"
	"testing"
)

func TestName_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := protobufBinding{}
	expected := "protobuf"
	actual := b.Name()
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
