package page

import (
	"fmt"
	"testing"
)

func TestGetPagePathBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := TargetPathDescriptor{}
	b := getPagePathBuilder(d)
	if b == nil {
		t.Errorf("getPagePathBuilder() = nil, want not nil")
	}
}
