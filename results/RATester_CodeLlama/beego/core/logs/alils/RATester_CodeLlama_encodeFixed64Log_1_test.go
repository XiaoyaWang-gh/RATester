package alils

import (
	"fmt"
	"testing"
)

func TestEncodeFixed64Log_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data := []byte{}
	offset := 0
	v := uint64(0)
	expected := 0
	actual := encodeFixed64Log(data, offset, v)
	if expected != actual {
		t.Errorf("encodeFixed64Log() expected %v, actual %v", expected, actual)
	}
}
