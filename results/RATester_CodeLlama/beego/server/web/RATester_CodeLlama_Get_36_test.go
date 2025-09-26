package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGet_36(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	c.Get()
	if c.Ctx.ResponseWriter.Status != http.StatusMethodNotAllowed {
		t.Errorf("Get() failed: want %d, got %d", http.StatusMethodNotAllowed, c.Ctx.ResponseWriter.Status)
	}
}
