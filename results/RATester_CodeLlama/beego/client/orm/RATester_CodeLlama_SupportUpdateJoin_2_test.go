package orm

import (
	"fmt"
	"testing"
)

func TestSupportUpdateJoin_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBase{}
	if !d.SupportUpdateJoin() {
		t.Errorf("TestSupportUpdateJoin failed")
	}
}
