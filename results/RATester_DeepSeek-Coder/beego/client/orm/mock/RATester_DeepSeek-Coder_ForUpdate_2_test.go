package mock

import (
	"fmt"
	"testing"
)

func TestForUpdate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	qs := d.ForUpdate()
	if _, ok := qs.(*DoNothingQuerySetter); !ok {
		t.Errorf("ForUpdate() did not return a DoNothingQuerySetter")
	}
}
