package orm

import (
	"fmt"
	"testing"
)

func TestSupportUpdateJoin_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBasePostgres{}
	expected := false
	result := db.SupportUpdateJoin()
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
