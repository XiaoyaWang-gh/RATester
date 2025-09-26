package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	apiextensionv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestCreatePluginMiddleware_1(t *testing.T) {
	type args struct {
		k8sClient Client
		ns        string
		plugins   map[string]apiextensionv1.JSON
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]dynamic.PluginConf
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

			got, err := createPluginMiddleware(tt.args.k8sClient, tt.args.ns, tt.args.plugins)
			if (err != nil) != tt.wantErr {
				t.Errorf("createPluginMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createPluginMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
