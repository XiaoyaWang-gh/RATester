package alils

import (
	"fmt"
	"testing"
)

func TestEncodeFixed32Log_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data := []byte{}
	offset := 0
	v := uint32(0)
	if got := encodeFixed32Log(data, offset, v); got != offset+4 {
		t.Errorf("encodeFixed32Log() = %v, want %v", got, offset+4)
	}
}
