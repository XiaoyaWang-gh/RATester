package utils

import (
	"bytes"
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
	args := []string{"text/plain"}
	a, err := e.Attach(r, filename, args...)
	if err != nil {
		t.Errorf("Attach() error = %v", err)
		return
	}
	if a.Filename != filename {
		t.Errorf("Attach() a.Filename = %v, want %v", a.Filename, filename)
	}
	if a.Header.Get("Content-Type") != "text/plain" {
		t.Errorf("Attach() a.Header.Get(\"Content-Type\") = %v, want %v", a.Header.Get("Content-Type"), "text/plain")
	}
	if a.Header.Get("Content-Disposition") != "attachment; filename=\"test.txt\"" {
		t.Errorf("Attach() a.Header.Get(\"Content-Disposition\") = %v, want %v", a.Header.Get("Content-Disposition"), "attachment; filename=\"test.txt\"")
	}
	if a.Header.Get("Content-Transfer-Encoding") != "base64" {
		t.Errorf("Attach() a.Header.Get(\"Content-Transfer-Encoding\") = %v, want %v", a.Header.Get("Content-Transfer-Encoding"), "base64")
	}
	if !bytes.Equal(a.Content, []byte("dGVzdA==")) {
		t.Errorf("Attach() a.Content = %v, want %v", a.Content, []byte("dGVzdA=="))
	}
}
