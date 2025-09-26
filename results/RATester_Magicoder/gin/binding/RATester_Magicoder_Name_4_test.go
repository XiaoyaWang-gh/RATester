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

	binding := protobufBinding{}
	name := binding.Name()
	if name != "protobuf" {
		t.Errorf("Expected 'protobuf', but got '%s'", name)
	}
}
