package orm

import (
	"fmt"
	"testing"
)

func TestForUpdate_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := &querySet{}
	qs.ForUpdate()
	if !qs.forUpdate {
		t.Error("ForUpdate() failed")
	}
}
