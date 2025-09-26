package orm

import (
	"fmt"
	"testing"
)

func TestClose_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &insertSet{}
	err := o.Close()
	if err != nil {
		t.Errorf("Close() error = %v", err)
		return
	}
}
