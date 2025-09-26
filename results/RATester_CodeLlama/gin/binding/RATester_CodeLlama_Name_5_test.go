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

	msgpackBinding := msgpackBinding{}
	if msgpackBinding.Name() != "msgpack" {
		t.Errorf("msgpackBinding.Name() = %v, want %v", msgpackBinding.Name(), "msgpack")
	}
}
