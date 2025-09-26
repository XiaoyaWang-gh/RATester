package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestAttach_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Email{}
	r := strings.NewReader("test")
	filename := "test.txt"
	args := []string{"text/plain", "testid"}
	a, err := e.Attach(r, filename, args...)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if a.Filename != filename {
		t.Errorf("Expected filename %s, got %s", filename, a.Filename)
	}
	if a.Header.Get("Content-Type") != "text/plain" {
		t.Errorf("Expected Content-Type text/plain, got %s", a.Header.Get("Content-Type"))
	}
	if a.Header.Get("Content-ID") != "<testid>" {
		t.Errorf("Expected Content-ID <testid>, got %s", a.Header.Get("Content-ID"))
	}
	if a.Header.Get("Content-Disposition") != "inline;\r\n filename=\"test.txt\"" {
		t.Errorf("Expected Content-Disposition inline;\r\n filename=\"test.txt\", got %s", a.Header.Get("Content-Disposition"))
	}
	if a.Header.Get("Content-Transfer-Encoding") != "base64" {
		t.Errorf("Expected Content-Transfer-Encoding base64, got %s", a.Header.Get("Content-Transfer-Encoding"))
	}
	if string(a.Content) != "dGVzdA==" {
		t.Errorf("Expected base64 encoded content, got %s", a.Content)
	}
}
