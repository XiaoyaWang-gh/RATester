package capture

import (
	"fmt"
	"net/http"
	"testing"
)

func TestStatus_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	crw := &captureResponseWriter{}
	crw.WriteHeader(http.StatusOK)
	if crw.Status() != http.StatusOK {
		t.Errorf("Status() = %d, want %d", crw.Status(), http.StatusOK)
	}
}
