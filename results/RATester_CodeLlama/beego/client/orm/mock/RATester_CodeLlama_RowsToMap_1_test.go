package mock

import (
	"fmt"
	"testing"
)

func TestRowsToMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	// TODO: setup mock
	// TODO: call RowsToMap
	// TODO: verify results
}
