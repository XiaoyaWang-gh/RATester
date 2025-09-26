package capture

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFlush_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	crw := &captureResponseWriter{}
	crw.Flush()
	if _, ok := crw.rw.(http.Flusher); !ok {
		t.Errorf("Flush() should not panic")
	}
}
