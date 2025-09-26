package transport

import (
	"fmt"
	"testing"
)

func TestNewRandomPrivateKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key, err := NewRandomPrivateKey()
	if err != nil {
		t.Errorf("NewRandomPrivateKey() failed, error: %v", err)
	}
	if key == nil {
		t.Errorf("NewRandomPrivateKey() failed, key is nil")
	}
}
