package acme

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-acme/lego/v4/lego"
)

func TestSetDefaults_2(t *testing.T) {
	tests := []struct {
		name string
		a    *Configuration
		want *Configuration
	}{
		{
			name: "default values",
			a:    &Configuration{},
			want: &Configuration{
				CAServer:             lego.LEDirectoryProduction,
				Storage:              "acme.json",
				KeyType:              "RSA4096",
				CertificatesDuration: 3 * 30 * 24,
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

			tt.a.SetDefaults()
			if !reflect.DeepEqual(tt.a, tt.want) {
				t.Errorf("SetDefaults() = %v, want %v", tt.a, tt.want)
			}
		})
	}
}
