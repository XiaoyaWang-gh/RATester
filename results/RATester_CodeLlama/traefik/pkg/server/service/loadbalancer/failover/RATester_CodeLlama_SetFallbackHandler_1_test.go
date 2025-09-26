package failover

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetFallbackHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &Failover{}
	handler := &http.ServeMux{}
	f.SetHandler(handler)
	f.SetFallbackHandler(handler)
	if f.fallbackHandler != handler {
		t.Errorf("expected %v, got %v", handler, f.fallbackHandler)
	}
}
