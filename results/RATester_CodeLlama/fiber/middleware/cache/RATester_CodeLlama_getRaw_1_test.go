package cache

import (
	"fmt"
	"testing"
)

func TestGetRaw_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	m := &manager{}
	key := "key"
	// act
	raw := m.getRaw(key)
	// assert
	if raw != nil {
		t.Errorf("getRaw() = %v, want nil", raw)
	}
}
