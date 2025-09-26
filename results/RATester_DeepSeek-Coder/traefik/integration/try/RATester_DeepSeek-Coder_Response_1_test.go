package try

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestResponse_1(t *testing.T) {
	t.Run("TestResponse", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		req, err := http.NewRequest("GET", "http://example.com", nil)
		if err != nil {
			t.Fatal(err)
		}

		timeout := 5 * time.Second

		resp, err := Response(req, timeout)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
		}
	})
}
