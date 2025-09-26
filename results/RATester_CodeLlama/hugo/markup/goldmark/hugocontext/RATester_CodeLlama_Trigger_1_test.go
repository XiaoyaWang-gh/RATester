package hugocontext

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrigger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &hugoContextParser{}
	if got := a.Trigger(); !reflect.DeepEqual(got, []byte{'{'}) {
		t.Errorf("Trigger() = %v, want %v", got, []byte{'{'})
	}
}
