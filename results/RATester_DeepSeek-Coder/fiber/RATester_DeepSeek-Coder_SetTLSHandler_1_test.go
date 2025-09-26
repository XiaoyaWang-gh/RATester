package fiber

import (
	"fmt"
	"testing"
)

func TestSetTLSHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	app := &App{}
	tlsHandler := &TLSHandler{}

	app.SetTLSHandler(tlsHandler)

	if app.tlsHandler != tlsHandler {
		t.Errorf("Expected tlsHandler to be %v, got %v", tlsHandler, app.tlsHandler)
	}
}
