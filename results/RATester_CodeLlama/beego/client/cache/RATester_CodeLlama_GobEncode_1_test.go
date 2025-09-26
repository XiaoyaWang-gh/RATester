package cache

import (
	"fmt"
	"testing"
)

func TestGobEncode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data := "hello world"
	buf, err := GobEncode(data)
	if err != nil {
		t.Errorf("GobEncode failed: %v", err)
	}
	if buf == nil {
		t.Errorf("GobEncode failed: buf is nil")
	}
}
