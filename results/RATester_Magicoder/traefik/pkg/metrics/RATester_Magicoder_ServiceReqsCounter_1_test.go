package metrics

import (
	"fmt"
	"testing"
)

func TestServiceReqsCounter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	registry := &standardRegistry{}

	counter := registry.ServiceReqsCounter()

	if counter == nil {
		t.Error("Expected a CounterWithHeaders, but got nil")
	}
}
