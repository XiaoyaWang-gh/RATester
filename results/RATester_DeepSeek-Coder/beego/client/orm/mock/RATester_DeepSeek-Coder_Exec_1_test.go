package mock

import (
	"fmt"
	"testing"
)

func TestExec_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingRawSetter{}
	_, err := d.Exec()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
