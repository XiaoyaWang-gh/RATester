package gin

import (
	"fmt"
	"testing"
)

func TestSetAccepted_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{}
	c.SetAccepted("json", "xml")
	if c.Accepted[0] != "json" || c.Accepted[1] != "xml" {
		t.Errorf("SetAccepted failed")
	}
}
