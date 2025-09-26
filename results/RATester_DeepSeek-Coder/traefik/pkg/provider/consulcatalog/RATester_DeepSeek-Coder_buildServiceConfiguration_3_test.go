package consulcatalog

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildServiceConfiguration_3(t *testing.T) {
	type args struct {
		item          itemData
		configuration *dynamic.HTTPConfiguration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{}
			if err := p.buildServiceConfiguration(tt.args.item, tt.args.configuration); (err != nil) != tt.wantErr {
				t.Errorf("Provider.buildServiceConfiguration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
