package orm

import (
	"fmt"
	"testing"
)

func TestNewOrmUsingDB_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var aliasName string
	if got := NewOrmUsingDB(aliasName); got != nil {
		t.Errorf("NewOrmUsingDB() = %v, want nil", got)
	}
}
