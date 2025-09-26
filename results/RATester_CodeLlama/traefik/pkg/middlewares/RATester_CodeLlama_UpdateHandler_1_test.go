package middlewares

import (
	"fmt"
	"net/http"
	"testing"
)

func TestUpdateHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HTTPHandlerSwitcher{}
	newHandler := &http.ServeMux{}
	h.UpdateHandler(newHandler)
	if h.handler.Get() != newHandler {
		t.Errorf("UpdateHandler failed")
	}
}
