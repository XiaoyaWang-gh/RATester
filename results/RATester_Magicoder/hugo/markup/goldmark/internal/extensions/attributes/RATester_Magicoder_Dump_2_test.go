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

	attrs := &attributesBlock{}
	source := []byte("test")
	level := 0
	attrs.Dump(source, level)
}
