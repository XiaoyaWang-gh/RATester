package plugins

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnzipModule_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	pName := "github.com/traefik/traefik"
	pVersion := "v3.0.0"
	c := &Client{
		HTTPClient: &http.Client{},
		baseURL:    &url.URL{},
		archives:   "",
		stateFile:  "",
		goPath:     "",
		sources:    "",
	}

	// when
	err := c.unzipModule(pName, pVersion)

	// then
	require.NoError(t, err)
}
