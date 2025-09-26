package mock

import (
	"fmt"
	"testing"
)

func TestExist_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	exist := d.Exist()
	if exist != true {
		t.Errorf("Expected true, got %v", exist)
	}
}
