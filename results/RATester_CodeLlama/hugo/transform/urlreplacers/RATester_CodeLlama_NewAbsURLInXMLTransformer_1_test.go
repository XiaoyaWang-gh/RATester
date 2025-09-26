package urlreplacers

import (
	"fmt"
	"testing"
)

func TestNewAbsURLInXMLTransformer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "path"
	transformer := NewAbsURLInXMLTransformer(path)
	if transformer == nil {
		t.Errorf("NewAbsURLInXMLTransformer(%s) = nil; want non-nil", path)
	}
}
