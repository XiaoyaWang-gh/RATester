package ingress

import (
	"fmt"
	"reflect"
	"testing"

	kinformers "k8s.io/client-go/informers"
	kclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func TestNewClientImpl_1(t *testing.T) {
	type args struct {
		clientset kclientset.Interface
	}
	tests := []struct {
		name string
		args args
		want *clientWrapper
	}{
		{
			name: "Test Case 1",
			args: args{
				clientset: &fake.Clientset{},
			},
			want: &clientWrapper{
				clientset:                   &fake.Clientset{},
				clusterScopeFactory:         nil,
				factoriesKube:               make(map[string]kinformers.SharedInformerFactory),
				factoriesSecret:             make(map[string]kinformers.SharedInformerFactory),
				factoriesIngress:            make(map[string]kinformers.SharedInformerFactory),
				ingressLabelSelector:        "",
				isNamespaceAll:              false,
				disableIngressClassInformer: false,
				disableClusterScopeInformer: false,
				watchedNamespaces:           nil,
				serverVersion:               nil,
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

			if got := newClientImpl(tt.args.clientset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newClientImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
