package utils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPrintPointerInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	buf := &bytes.Buffer{}
	headlen := 0
	pointers := &pointerInfo{
		prev: nil,
		n:    0,
		addr: 0,
		pos:  0,
		used: []int{},
	}

	PrintPointerInfo(buf, headlen, pointers)

	// Add assertions here to check the output of the function
}
