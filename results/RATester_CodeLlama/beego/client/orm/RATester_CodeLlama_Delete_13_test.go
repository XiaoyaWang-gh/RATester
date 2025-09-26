package orm

import (
	"fmt"
	"testing"
)

func TestDelete_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	o.Delete()
}
