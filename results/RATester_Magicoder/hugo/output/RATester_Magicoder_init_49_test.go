package output

import (
	"fmt"
	"sort"
	"testing"
)

func Testinit_49(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sort.Sort(DefaultFormats)

	// Add your test assertions here
}
