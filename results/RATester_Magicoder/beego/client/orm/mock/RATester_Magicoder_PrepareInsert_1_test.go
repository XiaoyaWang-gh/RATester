package mock

import (
	"fmt"
	"testing"
)

func TestPrepareInsert_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	_, err := d.PrepareInsert()
	if err != nil {
		t.Errorf("PrepareInsert() error = %v, wantErr %v", err, nil)
		return
	}
}
