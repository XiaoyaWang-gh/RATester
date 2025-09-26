package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetDefaultClientConf_2(t *testing.T) {
	tests := []struct {
		name string
		want ClientConfig
	}{
		{
			name: "Test case 1",
			want: ClientConfig{
				BaseConfig:       getDefaultBaseConf(),
				OidcClientConfig: getDefaultOidcClientConf(),
				TokenConfig:      getDefaultTokenConf(),
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

			if got := GetDefaultClientConf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDefaultClientConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
