package orm

import (
	"fmt"
	"testing"
)

func TestSupportUpdateJoin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseSqlite{}
	expected := false
	actual := db.SupportUpdateJoin()
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
