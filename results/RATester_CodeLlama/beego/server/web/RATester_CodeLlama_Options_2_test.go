package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestOptions_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	c.Options()
	if c.Ctx.ResponseWriter.Status != http.StatusMethodNotAllowed {
		t.Errorf("Options() failed")
	}
}
