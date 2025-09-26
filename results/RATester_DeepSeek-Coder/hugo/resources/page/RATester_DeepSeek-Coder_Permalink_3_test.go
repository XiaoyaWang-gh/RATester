package page

import (
	"fmt"
	"testing"
)

func TestPermalink_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := OutputFormat{
		permalink: "test-permalink",
	}

	result := o.Permalink()

	if result != "test-permalink" {
		t.Errorf("Expected 'test-permalink', got '%s'", result)
	}
}
