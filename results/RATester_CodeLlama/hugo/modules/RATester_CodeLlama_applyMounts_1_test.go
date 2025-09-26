package modules

import (
	"fmt"
	"testing"
)

func TestApplyMounts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &collector{}
	moduleImport := Import{
		Path: "github.com/gohugoio/hugo",
	}
	mod := &moduleAdapter{}

	err := c.applyMounts(moduleImport, mod)

	if err != nil {
		t.Fatal(err)
	}
}
