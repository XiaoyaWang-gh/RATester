package lazy

import (
	"fmt"
	"testing"
)

func TestShouldInitialize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ini := &Init{}
	if ini.shouldInitialize() {
		t.Errorf("shouldInitialize() = true, want false")
	}
}
