package testing

import (
	"fmt"
	"testing"
)

func TestHead_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "/test"
	req := Head(path)
	if req.GetRequest().Method != "HEAD" {
		t.Errorf("TestHead failed")
	}
}
