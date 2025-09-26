package commands

import (
	"errors"
	"fmt"
	"testing"
)

func TestBuildErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &hugoBuilderErrState{}
	e.setBuildErr(errors.New("test"))
	if e.buildErr() != errors.New("test") {
		t.Errorf("buildErr() = %v, want %v", e.buildErr(), errors.New("test"))
	}
}
