package urlreplacers

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestConsumeQuote_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &absurllexer{
		content: []byte("http://www.example.com/path/to/file.txt"),
		w:       ioutil.Discard,
		path:    []byte("."),
		quotes:  [][]byte{[]byte("\""), []byte("'")},
	}
	l.consumeQuote()
	if l.pos != 21 {
		t.Errorf("l.pos = %d, want %d", l.pos, 21)
	}
}
