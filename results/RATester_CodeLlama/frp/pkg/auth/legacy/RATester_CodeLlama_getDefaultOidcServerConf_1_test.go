package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetDefaultOidcServerConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	got := getDefaultOidcServerConf()
	want := OidcServerConfig{
		OidcIssuer:          "",
		OidcAudience:        "",
		OidcSkipExpiryCheck: false,
		OidcSkipIssuerCheck: false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getDefaultOidcServerConf() = %v, want %v", got, want)
	}
}
