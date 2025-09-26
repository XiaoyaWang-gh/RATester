package template

import (
	"fmt"
	"testing"
)

func TestTError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := context{}
	s := []byte("")
	c, _ = tError(c, s)
	if c.String() != "" {
		t.Errorf("tError(%v, %v) = %v, want %v", c, s, c, "")
	}
}
