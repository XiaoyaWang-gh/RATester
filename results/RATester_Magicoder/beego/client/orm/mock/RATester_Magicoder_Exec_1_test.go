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
	result, err := d.Exec()
	if err != nil {
		t.Errorf("Exec() error = %v, wantErr %v", err, false)
		return
	}
	if result != nil {
		t.Errorf("Exec() result = %v, want %v", result, nil)
	}
}
