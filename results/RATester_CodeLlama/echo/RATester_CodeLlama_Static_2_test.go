package echo

import (
	"fmt"
	"testing"
)

func TestStatic_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &Group{}
	pathPrefix := "pathPrefix"
	fsRoot := "fsRoot"
	g.Static(pathPrefix, fsRoot)
}
