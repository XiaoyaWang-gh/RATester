package file

import (
	"fmt"
	"testing"
)

func TestGetDb_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := GetDb()
	if db == nil {
		t.Errorf("GetDb() = %v; want not nil", db)
	}
}
