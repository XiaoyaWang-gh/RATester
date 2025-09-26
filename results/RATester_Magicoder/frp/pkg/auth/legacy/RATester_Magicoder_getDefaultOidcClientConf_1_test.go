package legacy

import (
	"fmt"
	"testing"
)

func TestGetDefaultOidcClientConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := getDefaultOidcClientConf()

	if conf.OidcClientID != "" {
		t.Errorf("Expected OidcClientID to be '', but got %s", conf.OidcClientID)
	}

	if conf.OidcClientSecret != "" {
		t.Errorf("Expected OidcClientSecret to be '', but got %s", conf.OidcClientSecret)
	}

	if conf.OidcAudience != "" {
		t.Errorf("Expected OidcAudience to be '', but got %s", conf.OidcAudience)
	}

	if conf.OidcScope != "" {
		t.Errorf("Expected OidcScope to be '', but got %s", conf.OidcScope)
	}

	if conf.OidcTokenEndpointURL != "" {
		t.Errorf("Expected OidcTokenEndpointURL to be '', but got %s", conf.OidcTokenEndpointURL)
	}

	if len(conf.OidcAdditionalEndpointParams) != 0 {
		t.Errorf("Expected OidcAdditionalEndpointParams to be empty, but got %v", conf.OidcAdditionalEndpointParams)
	}
}
