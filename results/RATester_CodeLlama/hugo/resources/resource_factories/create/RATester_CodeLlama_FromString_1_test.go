package create

import (
	"fmt"
	"testing"
)

func TestFromString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	r, err := c.FromString("targetPath", "content")
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("expected a resource")
	}
}
