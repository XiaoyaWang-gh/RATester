package resources

import (
	"fmt"
	"testing"
)

func TestHasTransformationPermalinkHash_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &resourceTransformations{}
	if r.hasTransformationPermalinkHash() {
		t.Errorf("hasTransformationPermalinkHash() = true, want false")
	}
}
