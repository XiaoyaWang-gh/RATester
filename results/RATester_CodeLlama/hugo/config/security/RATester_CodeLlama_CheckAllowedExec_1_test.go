package security

import (
	"fmt"
	"regexp"
	"testing"
)

func TestCheckAllowedExec_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := Config{
		Exec: Exec{
			Allow: Whitelist{
				patterns: []*regexp.Regexp{
					regexp.MustCompile("^/bin/sh$"),
					regexp.MustCompile("^/usr/bin/env$"),
				},
			},
		},
	}

	if err := c.CheckAllowedExec("/bin/sh"); err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if err := c.CheckAllowedExec("/usr/bin/env"); err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if err := c.CheckAllowedExec("/bin/bash"); err == nil {
		t.Fatalf("Expected error, got nil")
	}

	if err := c.CheckAllowedExec("/usr/bin/env"); err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
