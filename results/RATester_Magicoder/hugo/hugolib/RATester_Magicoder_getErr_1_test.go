package hugolib

import (
	"errors"
	"fmt"
	"testing"
)

func TestgetErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &fatalErrorHandler{
		err: errors.New("test error"),
	}

	got := f.getErr()
	want := errors.New("test error")

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
