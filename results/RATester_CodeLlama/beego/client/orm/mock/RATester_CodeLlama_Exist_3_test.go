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
	if d.Exist() != true {
		t.Errorf("Exist() = %v, want %v", d.Exist(), true)
	}
}
