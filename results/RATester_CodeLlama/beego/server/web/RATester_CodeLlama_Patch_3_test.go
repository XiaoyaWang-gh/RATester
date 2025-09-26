package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPatch_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	c.Patch()
	if c.Ctx.ResponseWriter.Status != http.StatusMethodNotAllowed {
		t.Errorf("Patch() failed")
	}
}
