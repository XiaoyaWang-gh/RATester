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

	b := &Bind{}
	b.Must()
	if b.should != false {
		t.Errorf("Must() should be false")
	}
}
