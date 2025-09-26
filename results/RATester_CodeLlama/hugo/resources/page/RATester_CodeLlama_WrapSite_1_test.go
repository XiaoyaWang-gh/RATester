package page

import (
	"fmt"
	"testing"
)

func TestWrapSite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var s Site
	if s == nil {
		t.Error("Site is nil")
	}
	if WrapSite(s) == nil {
		t.Error("WrapSite(s) is nil")
	}
}
