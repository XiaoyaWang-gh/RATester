package media

import (
	"fmt"
	"testing"
)

func TestInitMediaType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var m Type
	InitMediaType(&m)
	if m.IsZero() {
		t.Error("InitMediaType failed")
	}
}
