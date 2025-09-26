package security

import (
	"fmt"
	"regexp"
	"testing"
)

func TestCheckAllowedGetEnv_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := Config{
		Funcs: Funcs{
			Getenv: Whitelist{
				patterns: []*regexp.Regexp{
					regexp.MustCompile("^HUGO_"),
				},
			},
		},
	}

	if err := c.CheckAllowedGetEnv("HUGO_FOO"); err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if err := c.CheckAllowedGetEnv("FOO"); err == nil {
		t.Fatal("expected error")
	}
}
