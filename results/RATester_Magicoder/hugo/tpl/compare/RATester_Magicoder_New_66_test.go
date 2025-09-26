package compare

import (
	"fmt"
	"testing"
	"time"
)

func TestNew_66(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	loc, _ := time.LoadLocation("Local")
	ns := New(loc, true)

	if ns.loc != loc {
		t.Errorf("Expected loc to be %v, but got %v", loc, ns.loc)
	}

	if ns.caseInsensitive != true {
		t.Errorf("Expected caseInsensitive to be true, but got %v", ns.caseInsensitive)
	}
}
