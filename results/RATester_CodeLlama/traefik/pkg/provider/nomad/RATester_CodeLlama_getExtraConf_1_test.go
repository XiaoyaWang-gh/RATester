package nomad

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetExtraConf_1(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		p    *Provider
		args args
		want configuration
	}{
		{
			name: "test",
			p:    &Provider{},
			args: args{tags: []string{"traefik.enable=true", "traefik.nomad.canary=true"}},
			want: configuration{Enable: true, Canary: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.getExtraConf(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getExtraConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
