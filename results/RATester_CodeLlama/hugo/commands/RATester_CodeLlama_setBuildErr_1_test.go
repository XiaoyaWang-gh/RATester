package commands

import (
	"errors"
	"fmt"
	"testing"
)

func TestSetBuildErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &hugoBuilderErrState{}
	err := errors.New("test error")
	e.setBuildErr(err)
	if e.builderr != err {
		t.Errorf("expected %s, got %s", err, e.builderr)
	}
}
