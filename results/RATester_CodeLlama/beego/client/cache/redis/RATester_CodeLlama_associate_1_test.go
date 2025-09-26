package redis

import (
	"fmt"
	"testing"
)

func TestAssociate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rc := &Cache{}
	rc.key = "test"
	rc.skipEmptyPrefix = true
	originKey := "test"
	if rc.associate(originKey) != "test" {
		t.Errorf("Testassociate failed")
	}
}
