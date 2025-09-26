package hexec

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config/security"
)

func TestNpx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Exec{
		sc:          security.Config{},
		workingDir:  "/tmp",
		baseEnviron: []string{"PATH=/usr/local/bin:/usr/bin:/bin"},
	}

	name := "test"
	arg := []any{"arg1", "arg2"}

	_, err := e.npx(name, arg...)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
