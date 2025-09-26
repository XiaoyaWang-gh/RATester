package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetDefaultOidcClientConf_1(t *testing.T) {
	tests := []struct {
		name string
		want OidcClientConfig
	}{
		{
			name: "Test Case 1",
			want: OidcClientConfig{
				OidcClientID:                 "",
				OidcClientSecret:             "",
				OidcAudience:                 "",
				OidcScope:                    "",
				OidcTokenEndpointURL:         "",
				OidcAdditionalEndpointParams: make(map[string]string),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getDefaultOidcClientConf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDefaultOidcClientConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
