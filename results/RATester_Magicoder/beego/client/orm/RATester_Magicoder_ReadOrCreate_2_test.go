package orm

import (
	"fmt"
	"testing"
)

func TestReadOrCreate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &ormBase{}
	md := &struct{}{}
	col1 := "col1"
	cols := []string{"col2", "col3"}

	_, _, err := o.ReadOrCreate(md, col1, cols...)
	if err != nil {
		t.Errorf("ReadOrCreate returned an error: %v", err)
	}
}
