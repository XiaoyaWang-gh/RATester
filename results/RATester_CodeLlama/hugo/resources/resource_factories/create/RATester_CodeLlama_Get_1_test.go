package create

import (
	"fmt"
	"testing"
)

func TestGet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	pathname := "pathname"
	resource, err := c.Get(pathname)
	if err != nil {
		t.Fatal(err)
	}
	if resource == nil {
		t.Fatal("resource is nil")
	}
}
