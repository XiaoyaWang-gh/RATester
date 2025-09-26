package goldmark

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/goldmark/goldmark_config"
	"github.com/stretchr/testify/require"
)

func TestNewLinks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	cfg := goldmark_config.Config{}
	got := newLinks(cfg)
	require.NotNil(t, got)
	// TODO: assert
}
