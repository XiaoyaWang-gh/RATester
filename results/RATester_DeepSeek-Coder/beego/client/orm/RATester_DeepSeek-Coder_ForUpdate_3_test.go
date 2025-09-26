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

	o := querySet{}
	o.ForUpdate()
	if !o.forUpdate {
		t.Errorf("ForUpdate() did not set forUpdate to true")
	}
}
