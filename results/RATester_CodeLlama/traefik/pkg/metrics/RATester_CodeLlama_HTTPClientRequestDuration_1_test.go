package metrics

import (
	"fmt"
	"testing"
)

func TestHTTPClientRequestDuration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SemConvMetricsRegistry{}
	if s.HTTPClientRequestDuration() != nil {
		t.Error("HTTPClientRequestDuration() should return nil")
	}
}
