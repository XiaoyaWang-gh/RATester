package page

import (
	"fmt"
	"testing"
)

func TestNewPaginationURLFactory_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := TargetPathDescriptor{}
	newPaginationURLFactory(d)
}
