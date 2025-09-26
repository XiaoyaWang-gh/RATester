package gin

import (
	"fmt"
	"testing"
)

func TestBufApp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	buf := []byte{}
	s := "hello"
	w := 3
	c := byte('a')
	bufApp(&buf, s, w, c)
	if string(buf) != "hela" {
		t.Errorf("bufApp failed")
	}
}
