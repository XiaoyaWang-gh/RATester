package forwardedheaders

import (
	"fmt"
	"testing"
)

func TestSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := make(unsafeHeader)
	key := "testKey"
	value := "testValue"
	h.Set(key, value)

	if h[key][0] != value {
		t.Errorf("Expected %s, got %s", value, h[key][0])
	}
}
