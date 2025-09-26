package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewIngressRoutes_1(t *testing.T) {
	type args struct {
		c         *TraefikV1alpha1Client
		namespace string
	}
	tests := []struct {
		name string
		args args
		want *ingressRoutes
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

			if got := newIngressRoutes(tt.args.c, tt.args.namespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newIngressRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
