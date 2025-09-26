package media

import (
	"fmt"
	"testing"
)

func TestIsMarkdown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Type{
		SubType: "markdown",
	}
	if !m.IsMarkdown() {
		t.Errorf("IsMarkdown failed")
	}
}
