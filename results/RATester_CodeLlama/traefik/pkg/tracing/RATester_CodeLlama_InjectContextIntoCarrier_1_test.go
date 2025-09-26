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

	req := &http.Request{}
	InjectContextIntoCarrier(req)
	// TODO: Add test cases.
}
