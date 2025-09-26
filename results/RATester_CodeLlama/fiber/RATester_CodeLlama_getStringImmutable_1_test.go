package fiber

import (
	"fmt"
	"testing"
)

func TestGetStringImmutable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := []byte("hello")
	if getStringImmutable(b) != "hello" {
		t.Errorf("getStringImmutable() = %v, want %v", getStringImmutable(b), "hello")
	}
}
