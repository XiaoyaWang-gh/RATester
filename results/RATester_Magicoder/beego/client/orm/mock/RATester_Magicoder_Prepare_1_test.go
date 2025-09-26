package mock

import (
	"fmt"
	"testing"
)

func TestPrepare_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingRawSetter{}
	_, err := d.Prepare()
	if err != nil {
		t.Errorf("Prepare() failed, expected nil, got %v", err)
	}
}
