package hashing

import (
	"fmt"
	"testing"
)

func TestMD5FromStringHexEncoded_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := "hello"
	expected := "5d41402abc4b2a76b9719d911017c592"
	actual := MD5FromStringHexEncoded(f)
	if actual != expected {
		t.Errorf("MD5FromStringHexEncoded(%q) = %q; want %q", f, actual, expected)
	}
}
