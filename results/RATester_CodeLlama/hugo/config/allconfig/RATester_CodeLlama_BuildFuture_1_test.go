package allconfig

import (
	"fmt"
	"testing"
)

func TestBuildFuture_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.BuildFuture() {
		t.Error("BuildFuture() should be false")
	}
}
