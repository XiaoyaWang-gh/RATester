package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDelete_26(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	c.Delete()
	if c.Ctx.ResponseWriter.Status != http.StatusMethodNotAllowed {
		t.Errorf("Delete() = %v, want %v", c.Ctx.ResponseWriter.Status, http.StatusMethodNotAllowed)
	}
}
