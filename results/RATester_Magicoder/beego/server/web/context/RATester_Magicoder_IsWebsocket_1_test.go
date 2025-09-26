package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsWebsocket_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Header: http.Header{
					"Upgrade": []string{"websocket"},
				},
			},
		},
	}

	if !input.IsWebsocket() {
		t.Error("Expected IsWebsocket to return true")
	}
}
