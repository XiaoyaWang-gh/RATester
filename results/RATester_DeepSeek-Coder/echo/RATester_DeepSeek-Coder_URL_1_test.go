package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()

	h := func(c Context) error {
		return c.String(http.StatusOK, "test")
	}

	url := e.URL(h)
	if url != "/" {
		t.Errorf("Expected URL to be '/', got %s", url)
	}

	url = e.URL(h, "param")
	if url != "/param" {
		t.Errorf("Expected URL to be '/param', got %s", url)
	}
}
