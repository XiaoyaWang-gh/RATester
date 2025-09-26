package create

import (
	"fmt"
	"testing"
)

func TestBuildDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &contentBuilder{}
	b.buildDir()
}
