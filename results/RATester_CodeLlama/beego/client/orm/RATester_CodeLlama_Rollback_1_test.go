package orm

import (
	"fmt"
	"testing"
)

func TestRollback_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &filterOrmDecorator{}
	f.Rollback()
}
