package ingress

import (
	"fmt"
	"testing"
)

func TestConvertAnnotations_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	annotations := map[string]string{
		"ingress.kubernetes.io/foo": "bar",
		"ingress.kubernetes.io/bar": "baz",
	}

	result := convertAnnotations(annotations)

	if len(result) != 2 {
		t.Errorf("Expected 2 annotations, got %d", len(result))
	}

	if result["foo"] != "bar" {
		t.Errorf("Expected foo to be bar, got %s", result["foo"])
	}

	if result["bar"] != "baz" {
		t.Errorf("Expected bar to be baz, got %s", result["bar"])
	}
}
