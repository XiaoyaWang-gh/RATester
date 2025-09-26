package constraints

import (
	"fmt"
	"testing"
)

func TestOrTagFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tagFunc1 := func(tags []string) bool {
		return len(tags) > 0
	}

	tagFunc2 := func(tags []string) bool {
		return len(tags) > 1
	}

	orTaggedFunc := orTagFunc(tagFunc1, tagFunc2)

	tags := []string{"tag1", "tag2"}
	if !orTaggedFunc(tags) {
		t.Errorf("Expected true, got false")
	}

	tags = []string{"tag1"}
	if orTaggedFunc(tags) {
		t.Errorf("Expected false, got true")
	}

	tags = []string{}
	if orTaggedFunc(tags) {
		t.Errorf("Expected false, got true")
	}
}
