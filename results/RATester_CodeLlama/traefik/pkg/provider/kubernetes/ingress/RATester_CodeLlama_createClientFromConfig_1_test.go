package ingress

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func TestCreateClientFromConfig_1(t *testing.T) {
	type args struct {
		c *rest.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *clientWrapper
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				c: &rest.Config{
					Host: "127.0.0.1",
				},
			},
			want: &clientWrapper{
				clientset: kubernetes.NewForConfigOrDie(
					&rest.Config{
						Host: "127.0.0.1",
					},
				),
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

			got, err := createClientFromConfig(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("createClientFromConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createClientFromConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
