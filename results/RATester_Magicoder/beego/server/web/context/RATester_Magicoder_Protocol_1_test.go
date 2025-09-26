package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestProtocol_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Proto: "HTTP/1.1",
			},
		},
	}

	expected := "HTTP/1.1"
	actual := input.Protocol()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
