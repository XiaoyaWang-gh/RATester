package legacy

import (
	"fmt"
	"testing"
)

func TestGetDefaultOidcServerConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := OidcServerConfig{
		OidcIssuer:          "",
		OidcAudience:        "",
		OidcSkipExpiryCheck: false,
		OidcSkipIssuerCheck: false,
	}
	result := getDefaultOidcServerConf()
	if result != expected {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}
