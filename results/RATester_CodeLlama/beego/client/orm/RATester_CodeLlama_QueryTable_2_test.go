package orm

import (
	"fmt"
	"testing"
)

func TestQueryTable_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var (
		f                    *filterOrmDecorator
		ptrStructOrTableName interface{}
	)
	f.QueryTable(ptrStructOrTableName)
}
