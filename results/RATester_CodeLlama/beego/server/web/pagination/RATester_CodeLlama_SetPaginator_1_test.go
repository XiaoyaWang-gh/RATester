package pagination

import (
	"fmt"
	"testing"
)

func TestSetPaginator_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
