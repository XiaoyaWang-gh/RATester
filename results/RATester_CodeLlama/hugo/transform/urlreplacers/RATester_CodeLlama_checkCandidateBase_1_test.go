package urlreplacers

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCheckCandidateBase_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &absurllexer{
		content: []byte("https://www.google.com/search?q=golang"),
		w:       ioutil.Discard,
		path:    []byte("."),
	}
	checkCandidateBase(l)
	if l.pos != 22 {
		t.Errorf("Expected pos to be 22, got %d", l.pos)
	}
	if l.start != 22 {
		t.Errorf("Expected start to be 22, got %d", l.start)
	}
}
