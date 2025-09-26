package fiber

import (
	"fmt"
	"testing"
)

func TestShould_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Bind{}
	b.Should()
	if b.should != true {
		t.Errorf("Should() failed")
	}
}
