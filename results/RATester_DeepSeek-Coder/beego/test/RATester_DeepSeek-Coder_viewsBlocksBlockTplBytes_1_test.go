package test

import (
	"fmt"
	"testing"
)

func TestViewsBlocksBlockTplBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data, err := viewsBlocksBlockTplBytes()
	if err != nil {
		t.Errorf("viewsBlocksBlockTplBytes() error = %v", err)
		return
	}
	if len(data) == 0 {
		t.Errorf("viewsBlocksBlockTplBytes() expected data, got nothing")
	}
}
