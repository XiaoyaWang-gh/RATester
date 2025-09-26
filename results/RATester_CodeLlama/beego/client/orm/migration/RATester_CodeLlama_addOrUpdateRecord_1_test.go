package migration

import (
	"fmt"
	"testing"
)

func TestAddOrUpdateRecord_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	m := &Migration{}
	name := "test"
	status := "test"
	err := m.addOrUpdateRecord(name, status)
	if err != nil {
		t.Errorf("addOrUpdateRecord() error = %v", err)
		return
	}
	// TODO: test cases
	// TODO: teardown test
}
