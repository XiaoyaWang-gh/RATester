package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestNewIngressRouteTCPs_1(t *testing.T) {
	type args struct {
		c         *TraefikV1alpha1Client
		namespace string
	}
	tests := []struct {
		name string
		args args
		want *ingressRouteTCPs
	}{
		{
			name: "Test Case 1",
			args: args{
				c: &TraefikV1alpha1Client{
					restClient: &rest.RESTClient{},
				},
				namespace: "default",
			},
			want: &ingressRouteTCPs{
				client: &rest.RESTClient{},
				ns:     "default",
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

			if got := newIngressRouteTCPs(tt.args.c, tt.args.namespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newIngressRouteTCPs() = %v, want %v", got, tt.want)
			}
		})
	}
}
