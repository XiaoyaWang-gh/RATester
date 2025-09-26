package berror

import (
	"fmt"
	"testing"
)

func TestDesc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cd := &codeDefinition{
		code:   1,
		module: "test",
		desc:   "test desc",
		name:   "test name",
	}

	got := cd.Desc()
	want := "test desc"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
