package gin

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteHeaderNow_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &responseWriter{
		ResponseWriter: httptest.NewRecorder(),
		status:         200,
	}
	w.WriteHeaderNow()
	if w.status != 200 {
		t.Errorf("w.status = %d; want 200", w.status)
	}
}
