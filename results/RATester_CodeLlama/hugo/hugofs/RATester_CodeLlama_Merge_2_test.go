package hugofs

import (
	"fmt"
	"testing"
)

func TestMerge_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &FileMeta{}
	from := &FileMeta{}
	m.Merge(from)
}
