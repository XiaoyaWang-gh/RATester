package urlreplacers

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPosAfterURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &absurllexer{
		content: []byte("http://www.example.com/path/to/file.html"),
		w:       ioutil.Discard,
		path:    []byte("."),
	}
	q := []byte("http://www.example.com/path/to/file.html")
	if l.posAfterURL(q) != 0 {
		t.Errorf("posAfterURL(%q) = %d, want 0", q, l.posAfterURL(q))
	}
}
