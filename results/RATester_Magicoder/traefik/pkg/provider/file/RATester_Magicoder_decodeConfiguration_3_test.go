package file

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestdecodeConfiguration_3(t *testing.T) {
	type args struct {
		filePath string
		content  string
	}
	tests := []struct {
		name    string
		args    args
		want    *dynamic.Configuration
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
			got, err := p.decodeConfiguration(tt.args.filePath, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.decodeConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.decodeConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}
