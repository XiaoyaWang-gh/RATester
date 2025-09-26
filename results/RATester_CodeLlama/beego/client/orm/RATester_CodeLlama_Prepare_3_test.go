package orm

import (
	"fmt"
	"testing"
)

func TestPrepare_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &rawSet{}
	_, err := o.Prepare()
	if err != nil {
		t.Errorf("Prepare() error = %v", err)
		return
	}
}
