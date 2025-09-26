package connectionheader

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRemove_1(t *testing.T) {
	t.Run("TestRemove", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		req, err := http.NewRequest("GET", "http://example.com", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Connection", "upgrade")
		req.Header.Set("Upgrade", "websocket")

		newReq := Remove(req)

		if newReq.Header.Get("Connection") != "" {
			t.Errorf("Expected Connection header to be removed, got %s", newReq.Header.Get("Connection"))
		}

		if newReq.Header.Get("Upgrade") != "websocket" {
			t.Errorf("Expected Upgrade header to be preserved, got %s", newReq.Header.Get("Upgrade"))
		}
	})
}
