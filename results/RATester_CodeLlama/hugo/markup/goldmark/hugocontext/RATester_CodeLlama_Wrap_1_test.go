package hugocontext

import (
	"fmt"
	"testing"
)

func TestWrap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := []byte("hello world")
	pid := uint64(123456789)
	expected := "prefix pid=123456789\nhello worldprefix)\n"
	actual := Wrap(b, pid)
	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}
