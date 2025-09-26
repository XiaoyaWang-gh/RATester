package observability

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHijack_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	recorder := &statusCodeRecorder{
		status: http.StatusOK,
	}

	_, _, err := recorder.Hijack()
	if err != nil {
		t.Errorf("Hijack() returned an error: %v", err)
	}
}
