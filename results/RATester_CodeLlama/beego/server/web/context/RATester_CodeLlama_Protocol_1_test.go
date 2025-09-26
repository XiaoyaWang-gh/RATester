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
	if input.Protocol() != "HTTP/1.1" {
		t.Errorf("Protocol() = %v, want HTTP/1.1", input.Protocol())
	}
}
