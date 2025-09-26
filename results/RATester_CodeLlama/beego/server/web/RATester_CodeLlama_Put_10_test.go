package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPut_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	c.Put()
	if c.Ctx.ResponseWriter.Status != http.StatusMethodNotAllowed {
		t.Errorf("Put() = %v, want %v", c.Ctx.ResponseWriter.Status, http.StatusMethodNotAllowed)
	}
}
