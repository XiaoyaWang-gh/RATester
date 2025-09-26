package commands

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config/allconfig"
)

func TestCopyStatic_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &hugoBuilder{
		r: &rootCommand{
			source: "/path/to/source",
		},
		conf: &commonConfig{
			configs: &allconfig.Configs{},
		},
	}

	_, err := h.copyStatic()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
