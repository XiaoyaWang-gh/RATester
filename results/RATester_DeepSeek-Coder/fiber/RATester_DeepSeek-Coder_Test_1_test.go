package fiber

import (
	"fmt"
	"net/http"
	"testing"
)

func TestTest_1(t *testing.T) {
	t.Run("Test", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		app := &App{}
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		res, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
		}
	})
}
