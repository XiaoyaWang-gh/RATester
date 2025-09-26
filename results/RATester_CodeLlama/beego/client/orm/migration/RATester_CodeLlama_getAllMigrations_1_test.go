package migration

import (
	"fmt"
	"testing"
)

func TestGetAllMigrations_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	migs, err := getAllMigrations()
	if err != nil {
		t.Errorf("getAllMigrations error: %v", err)
	}
	if len(migs) == 0 {
		t.Errorf("getAllMigrations error: %v", migs)
	}
}
