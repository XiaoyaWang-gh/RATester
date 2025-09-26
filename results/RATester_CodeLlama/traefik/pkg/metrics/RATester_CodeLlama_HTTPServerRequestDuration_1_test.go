package metrics

import (
	"fmt"
	"testing"
)

func TestHTTPServerRequestDuration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SemConvMetricsRegistry{}
	if s.HTTPServerRequestDuration() != nil {
		t.Error("expected nil")
	}
}
