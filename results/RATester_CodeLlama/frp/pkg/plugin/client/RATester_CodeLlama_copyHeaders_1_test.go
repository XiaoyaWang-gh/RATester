package client

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCopyHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dst := make(http.Header)
	src := make(http.Header)
	src.Add("key1", "value1")
	src.Add("key1", "value2")
	src.Add("key2", "value3")
	copyHeaders(dst, src)
	if len(dst) != 2 {
		t.Errorf("dst len = %d, want 2", len(dst))
	}
	if dst.Get("key1") != "value1,value2" {
		t.Errorf("dst.Get(key1) = %q, want value1,value2", dst.Get("key1"))
	}
	if dst.Get("key2") != "value3" {
		t.Errorf("dst.Get(key2) = %q, want value3", dst.Get("key2"))
	}
}
