package jsconfig

import (
	"fmt"
	"testing"
)

func TestAddSourceRoot_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Builder{}
	root := "root"
	b.AddSourceRoot(root)
	if _, found := b.sourceRoots[root]; !found {
		t.Errorf("expected %q to be added to sourceRoots", root)
	}
}
