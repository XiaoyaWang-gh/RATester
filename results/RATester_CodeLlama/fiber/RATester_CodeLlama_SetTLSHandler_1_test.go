package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestSetTLSHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	tlsHandler := &TLSHandler{}
	app.SetTLSHandler(tlsHandler)
	assert.Equal(t, tlsHandler, app.tlsHandler)
}
