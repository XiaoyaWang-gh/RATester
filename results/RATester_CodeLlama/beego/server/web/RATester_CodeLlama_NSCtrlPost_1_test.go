package web

import (
	"fmt"
	"testing"
)

func TestNSCtrlPost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var rootpath string
	var f interface{}
	var link LinkNamespace
	link = NSCtrlPost(rootpath, f)
	if link == nil {
		t.Errorf("NSCtrlPost() = %v, want %v", link, nil)
	}
}
