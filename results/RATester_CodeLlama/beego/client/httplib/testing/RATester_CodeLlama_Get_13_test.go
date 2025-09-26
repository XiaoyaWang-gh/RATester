package testing

import (
	"fmt"
	"testing"
)

func TestGet_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "/test"
	req := Get(path)
	if req.GetRequest().URL.Path != path {
		t.Errorf("TestGet failed")
	}
}
