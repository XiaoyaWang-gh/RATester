package page

import (
	"fmt"
	"testing"
)

func TestSocial_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := testSite{}
	if s.Social() == nil {
		t.Errorf("Social() should not be nil")
	}
}
