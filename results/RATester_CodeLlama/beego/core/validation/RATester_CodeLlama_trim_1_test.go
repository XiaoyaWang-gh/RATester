package validation

import (
	"fmt"
	"testing"
)

func TestTrim_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	key := "test"
	s := []string{"test"}
	ts, err := trim(name, key, s)
	if err != nil {
		t.Errorf("trim error: %v", err)
	}
	if len(ts) != 2 {
		t.Errorf("trim error: %v", ts)
	}
}
