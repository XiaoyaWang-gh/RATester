package file

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildConfiguration_5(t *testing.T) {
	type args struct {
		p *Provider
	}
	tests := []struct {
		name    string
		args    args
		want    *dynamic.Configuration
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				p: &Provider{
					Directory: "test",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.args.p.buildConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("buildConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildConfiguration() got = %v, want %v", got, tt.want)
			}
		})
	}
}
