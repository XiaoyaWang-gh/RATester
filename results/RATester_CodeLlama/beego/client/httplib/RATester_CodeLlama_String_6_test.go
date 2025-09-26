package httplib

import (
	"fmt"
	"testing"
)

func TestString_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{}
	b.body = []byte("test")
	str, err := b.String()
	if err != nil {
		t.Error(err)
	}
	if str != "test" {
		t.Error("string not equal")
	}
}
