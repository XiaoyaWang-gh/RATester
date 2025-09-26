package config

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
)

func TestMatchHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		Headers: []Headers{
			{
				Values: map[string]any{
					"Content-Type":  "text/html",
					"Cache-Control": "no-cache",
				},
			},
			{
				Values: map[string]any{
					"Content-Type":  "application/json",
					"Cache-Control": "max-age=3600",
				},
			},
		},
		compiledHeaders: []glob.Glob{
			glob.MustCompile("*"),
			glob.MustCompile("*"),
		},
	}

	matches := server.MatchHeaders("*")

	if len(matches) != 4 {
		t.Errorf("Expected 4 matches, got %d", len(matches))
	}

	for _, match := range matches {
		if match.Key != "Cache-Control" && match.Key != "Content-Type" {
			t.Errorf("Unexpected key: %s", match.Key)
		}
		if match.Value != "max-age=3600" && match.Value != "no-cache" && match.Value != "text/html" && match.Value != "application/json" {
			t.Errorf("Unexpected value: %s", match.Value)
		}
	}
}
