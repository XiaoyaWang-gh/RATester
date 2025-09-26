package attributes

import (
	"fmt"
	"testing"
)

func TestDump_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	source := []byte("")
	level := 0
	a := &attributesBlock{}
	a.Dump(source, level)
}
