package config

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
)

func TestMatchRedirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Server{}
	s.compiledRedirects = []glob.Glob{glob.MustCompile("a")}
	s.Redirects = []Redirect{{To: "b"}}

	if got := s.MatchRedirect("a"); got.To != "b" {
		t.Errorf("MatchRedirect() = %v, want %v", got, "b")
	}
}
