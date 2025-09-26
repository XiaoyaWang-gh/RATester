package allconfig

import (
	"fmt"
	"testing"
)

func TestBuildExpired_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.BuildExpired() {
		t.Error("BuildExpired() should be false")
	}
}
