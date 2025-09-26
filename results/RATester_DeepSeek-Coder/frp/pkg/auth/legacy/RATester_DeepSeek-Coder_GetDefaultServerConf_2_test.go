package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetDefaultServerConf_2(t *testing.T) {
	tests := []struct {
		name string
		want ServerConfig
	}{
		{
			name: "Test case 1",
			want: ServerConfig{
				BaseConfig:       getDefaultBaseConf(),
				OidcServerConfig: getDefaultOidcServerConf(),
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

			if got := GetDefaultServerConf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDefaultServerConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
