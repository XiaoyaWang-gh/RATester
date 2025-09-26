package binding

import (
	"fmt"
	"testing"
)

func TestName_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	uri := uriBinding{}
	name := uri.Name()
	if name != "uri" {
		t.Errorf("Expected 'uri', but got '%s'", name)
	}
}
