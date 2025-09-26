package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestMiddlewares_4(t *testing.T) {
	type fields struct {
		restClient rest.Interface
	}
	type args struct {
		namespace string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MiddlewareInterface
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

			c := &TraefikV1alpha1Client{
				restClient: tt.fields.restClient,
			}
			if got := c.Middlewares(tt.args.namespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middlewares() = %v, want %v", got, tt.want)
			}
		})
	}
}
