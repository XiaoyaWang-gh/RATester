package hugolib

import (
	"fmt"
	"testing"
)

func TestNewContentFactory_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoSites{}
	f := NewContentFactory(h)
	if f.h != h {
		t.Errorf("Expected hugo sites to be set")
	}
}
