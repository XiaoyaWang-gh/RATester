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
			name: "Test newClientImpl",
			args: args{
				clientset: fake.NewSimpleClientset(),
			},
			want: &clientWrapper{
				clientset:        fake.NewSimpleClientset(),
				factoriesKube:    make(map[string]kinformers.SharedInformerFactory),
				factoriesSecret:  make(map[string]kinformers.SharedInformerFactory),
				factoriesIngress: make(map[string]kinformers.SharedInformerFactory),
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

			got := newClientImpl(tt.args.clientset)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newClientImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
