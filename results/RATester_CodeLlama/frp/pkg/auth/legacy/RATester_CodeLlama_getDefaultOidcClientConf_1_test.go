package legacy

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetDefaultOidcClientConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := getDefaultOidcClientConf()
	assert.Equal(t, "", conf.OidcClientID)
	assert.Equal(t, "", conf.OidcClientSecret)
	assert.Equal(t, "", conf.OidcAudience)
	assert.Equal(t, "", conf.OidcScope)
	assert.Equal(t, "", conf.OidcTokenEndpointURL)
	assert.Equal(t, make(map[string]string), conf.OidcAdditionalEndpointParams)
}
