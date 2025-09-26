package urlreplacers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCheckCandidateSrcset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &absurllexer{
		content: []byte("srcset"),
		w:       &bytes.Buffer{},
		path:    []byte("."),
	}
	checkCandidateSrcset(l)
	if l.w.(*bytes.Buffer).String() != "srcset" {
		t.Errorf("Expected %s, got %s", "srcset", l.w.(*bytes.Buffer).String())
	}
}
