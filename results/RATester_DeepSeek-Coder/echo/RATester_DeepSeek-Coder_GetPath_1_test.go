package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetPath_1(t *testing.T) {
	t.Run("TestGetPath", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		req, err := http.NewRequest("GET", "/test/path?test=1", nil)
		if err != nil {
			t.Fatal(err)
		}

		got := GetPath(req)
		want := "/test/path"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
