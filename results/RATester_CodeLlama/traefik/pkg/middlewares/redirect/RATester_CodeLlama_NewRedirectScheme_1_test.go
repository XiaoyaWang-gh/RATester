package redirect

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewRedirectScheme_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	ctx := context.Background()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	conf := dynamic.RedirectScheme{
		Scheme: "https",
		Port:   "443",
	}
	name := "test"

	// act
	handler, err := NewRedirectScheme(ctx, next, conf, name)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, handler)
}
