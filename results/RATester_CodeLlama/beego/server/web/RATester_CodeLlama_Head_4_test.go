package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHead_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	c.Head()
	if c.Ctx.ResponseWriter.Status != http.StatusMethodNotAllowed {
		t.Errorf("Head() failed")
	}
}
