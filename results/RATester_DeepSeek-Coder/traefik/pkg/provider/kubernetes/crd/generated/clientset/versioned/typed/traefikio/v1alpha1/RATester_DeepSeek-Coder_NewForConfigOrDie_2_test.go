package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestNewForConfigOrDie_2(t *testing.T) {
	type args struct {
		c *rest.Config
	}
	tests := []struct {
		name string
		args args
		want *TraefikV1alpha1Client
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

			if got := NewForConfigOrDie(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewForConfigOrDie() = %v, want %v", got, tt.want)
			}
		})
	}
}
