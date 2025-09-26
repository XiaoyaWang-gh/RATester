package common

import (
	"fmt"
	"testing"
)

func TestPutBufPoolMax_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	buf := make([]byte, 1024)
	PutBufPoolMax(buf)
	if buf != nil {
		t.Errorf("buf should be nil")
	}
}
