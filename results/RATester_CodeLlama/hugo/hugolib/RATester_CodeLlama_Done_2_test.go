package hugolib

import (
	"fmt"
	"testing"
)

func TestDone_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &fatalErrorHandler{}
	f.donec = make(chan bool)
	if f.Done() != f.donec {
		t.Errorf("f.Done() != f.donec")
	}
}
