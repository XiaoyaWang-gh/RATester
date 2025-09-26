package urlreplacers

import (
	"fmt"
	"testing"
)

func TestNewAbsURLTransformer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "path"
	transformer := NewAbsURLTransformer(path)
	if transformer == nil {
		t.Errorf("NewAbsURLTransformer(%s) = nil; want non-nil", path)
	}
}
