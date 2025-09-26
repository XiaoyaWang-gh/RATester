package page

import (
	"fmt"
	"testing"
)

func TestRelPermalink_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := OutputFormat{
		relPermalink: "test",
	}
	if o.RelPermalink() != "test" {
		t.Errorf("Expected %q, got %q", "test", o.RelPermalink())
	}
}
