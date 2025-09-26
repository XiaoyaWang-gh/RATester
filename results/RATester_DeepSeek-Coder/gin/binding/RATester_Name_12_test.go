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
	expected := "uri"
	actual := uri.Name()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
