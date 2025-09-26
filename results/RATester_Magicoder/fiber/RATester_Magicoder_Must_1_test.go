package fiber

import (
	"fmt"
	"testing"
)

func TestMust_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Bind{
		should: true,
	}

	b.Must()

	if b.should != false {
		t.Errorf("Expected should to be false, but got %v", b.should)
	}
}
