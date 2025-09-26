package gin

import (
	"fmt"
	"testing"
)

func TestRun_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		RedirectTrailingSlash: true,
	}

	err := engine.Run(":8080")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
