package orm

import (
	"fmt"
	"testing"
)

func TestExist_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := querySet{}
	exist := qs.Exist()
	if exist != true {
		t.Errorf("Expected true, got %v", exist)
	}
}
