package cors

import (
	"fmt"
	"testing"
)

func TestIsOriginAllowed_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &Options{
		AllowAllOrigins: true,
	}
	if !o.IsOriginAllowed("*") {
		t.Errorf("IsOriginAllowed() = %v, want %v", false, true)
	}
}
