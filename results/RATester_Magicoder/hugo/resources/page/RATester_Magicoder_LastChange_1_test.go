package page

import (
	"fmt"
	"testing"
)

func TestLastChange_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testSite := testSite{}
	lastChange := testSite.LastChange()

	if lastChange.IsZero() {
		t.Errorf("LastChange() should not return a zero time")
	}
}
