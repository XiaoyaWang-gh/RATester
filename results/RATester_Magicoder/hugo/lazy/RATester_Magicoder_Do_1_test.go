package lazy

import (
	"fmt"
	"testing"
)

func TestDo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &onceMore{}

	// Test case 1: Function is executed once
	executed := false
	o.Do(func() {
		executed = true
	})
	if !executed {
		t.Error("Function was not executed")
	}
	if !o.Done() {
		t.Error("OnceMore is not done")
	}
	if o.InProgress() {
		t.Error("OnceMore is in progress")
	}

	// Test case 2: Function is not executed again
	executed = false
	o.Do(func() {
		executed = true
	})
	if executed {
		t.Error("Function was executed again")
	}
	if !o.Done() {
		t.Error("OnceMore is not done")
	}
	if o.InProgress() {
		t.Error("OnceMore is in progress")
	}

	// Test case 3: Function is executed again after reset
	o.ResetWithLock()
	executed = false
	o.Do(func() {
		executed = true
	})
	if !executed {
		t.Error("Function was not executed again")
	}
	if !o.Done() {
		t.Error("OnceMore is not done")
	}
	if o.InProgress() {
		t.Error("OnceMore is in progress")
	}
}
