package tracing

import (
	"fmt"
	"net/http"
	"testing"
)

func TestInjectContextIntoCarrier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &http.Request{
		Header: make(http.Header),
	}

	InjectContextIntoCarrier(req)

	// Add assertions here
}
