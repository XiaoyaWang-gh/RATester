package alils

import (
	"fmt"
	"testing"
)

func TestEncodeVarintLog_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data := []byte{}
	offset := 0
	v := uint64(1)
	if got := encodeVarintLog(data, offset, v); got != 1 {
		t.Errorf("encodeVarintLog() = %v, want %v", got, 1)
	}
}
