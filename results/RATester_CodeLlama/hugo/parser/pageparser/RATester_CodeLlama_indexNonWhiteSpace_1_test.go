package pageparser

import (
	"fmt"
	"testing"
)

func TestIndexNonWhiteSpace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := []byte("hello world")
	in := 'w'
	if indexNonWhiteSpace(s, in) != -1 {
		t.Errorf("indexNonWhiteSpace(%q, %q) = %d; want -1", s, in, indexNonWhiteSpace(s, in))
	}
}
