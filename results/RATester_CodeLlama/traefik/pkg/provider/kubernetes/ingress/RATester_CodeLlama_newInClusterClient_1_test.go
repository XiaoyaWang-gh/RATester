package ingress

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewInClusterClient_1(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name    string
		args    args
		want    *clientWrapper
		wantErr bool
	}{
		{
			name: "test newInClusterClient",
			args: args{
				endpoint: "127.0.0.1:8080",
			},
			want: &clientWrapper{
				clientset:                   nil,
				clusterScopeFactory:         nil,
				factoriesKube:               nil,
				factoriesSecret:             nil,
				factoriesIngress:            nil,
				ingressLabelSelector:        "",
				isNamespaceAll:              false,
				disableIngressClassInformer: false,
				disableClusterScopeInformer: false,
				watchedNamespaces:           nil,
				serverVersion:               nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := newInClusterClient(tt.args.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("newInClusterClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInClusterClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
