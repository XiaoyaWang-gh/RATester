package hexec

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/config/security"
)

func Testnpx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Exec{
		sc:         security.Config{},
		workingDir: "/path/to/working/dir",
		baseEnviron: []string{
			"ENV1=value1",
			"ENV2=value2",
		},
		npxInit:      sync.Once{},
		npxAvailable: true,
	}

	name := "test-name"
	arg := []any{"arg1", "arg2"}

	expectedRunner, expectedError := e.new(name, name, arg...)

	runner, err := e.npx(name, arg...)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if runner != expectedRunner {
		t.Errorf("Expected runner %v, but got %v", expectedRunner, runner)
	}
}
