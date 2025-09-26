package orm

import (
	"fmt"
	"testing"
)

func TestReplaceMarks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d dbBase
	var query *string
	d.ReplaceMarks(query)
}
