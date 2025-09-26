package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/source"
)

func TestreadData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoSites{}
	f := &source.File{}

	_, err := h.readData(f)
	if err != nil {
		t.Errorf("readData returned an error: %v", err)
	}
}
