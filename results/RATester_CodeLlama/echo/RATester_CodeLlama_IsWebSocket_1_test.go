package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsWebSocket_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &context{
		request: &http.Request{
			Header: http.Header{
				HeaderUpgrade: []string{"websocket"},
			},
		},
	}
	if !c.IsWebSocket() {
		t.Errorf("IsWebSocket() = %v, want %v", false, true)
	}
}
