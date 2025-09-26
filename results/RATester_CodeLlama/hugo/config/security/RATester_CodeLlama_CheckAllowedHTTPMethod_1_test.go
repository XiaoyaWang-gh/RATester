package security

import (
	"fmt"
	"regexp"
	"testing"
)

func TestCheckAllowedHTTPMethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := Config{
		HTTP: HTTP{
			Methods: Whitelist{
				patterns: []*regexp.Regexp{
					regexp.MustCompile("GET"),
					regexp.MustCompile("POST"),
				},
			},
		},
	}

	if err := c.CheckAllowedHTTPMethod("GET"); err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if err := c.CheckAllowedHTTPMethod("POST"); err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if err := c.CheckAllowedHTTPMethod("PUT"); err == nil {
		t.Fatalf("expected error, got nil")
	}
}
