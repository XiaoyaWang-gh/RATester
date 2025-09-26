package orm

import (
	"fmt"
	"testing"
)

func TestnewdbBaseTidb_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := newdbBaseTidb()
	if db == nil {
		t.Error("newdbBaseTidb() failed")
	}
}
