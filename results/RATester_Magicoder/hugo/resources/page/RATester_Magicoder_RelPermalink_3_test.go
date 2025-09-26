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
		Rel:          "canonical",
		relPermalink: "/canonical",
	}

	if o.RelPermalink() != "/canonical" {
		t.Errorf("Expected RelPermalink to be '/canonical', but got %s", o.RelPermalink())
	}
}
