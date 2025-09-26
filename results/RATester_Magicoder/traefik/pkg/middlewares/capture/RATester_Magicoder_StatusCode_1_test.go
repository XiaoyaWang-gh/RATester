package capture

import (
	"fmt"
	"net/http"
	"testing"
)

func TestStatusCode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	capture := &Capture{
		crw: &captureResponseWriter{
			status: http.StatusOK,
		},
	}

	expected := http.StatusOK
	actual := capture.StatusCode()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
