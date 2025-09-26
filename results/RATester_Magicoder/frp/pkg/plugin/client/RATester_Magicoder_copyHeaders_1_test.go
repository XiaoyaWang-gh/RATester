package client

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCopyHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	src := http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer token"},
	}
	dst := http.Header{}

	copyHeaders(dst, src)

	if len(dst) != len(src) {
		t.Errorf("Expected %d headers, got %d", len(src), len(dst))
	}

	for key, values := range src {
		if dstValues := dst.Values(key); !reflect.DeepEqual(values, dstValues) {
			t.Errorf("Expected %v for key %s, got %v", values, key, dstValues)
		}
	}
}
