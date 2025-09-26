package captcha

import (
	"fmt"
	"testing"
)

func TestUint32_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &siprng{}
	p.rekey()
	for i := 0; i < 1000; i++ {
		if p.Uint32() != uint32(p.Uint64()) {
			t.Fatalf("Uint32() = %d, want %d", p.Uint32(), uint32(p.Uint64()))
		}
	}
}
