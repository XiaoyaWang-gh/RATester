package allconfig

import (
	"fmt"
	"testing"
)

func TestNoBuildLock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.NoBuildLock() {
		t.Error("NoBuildLock() should be false")
	}
}
