package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewServersTransports_1(t *testing.T) {
	type args struct {
		c         *TraefikV1alpha1Client
		namespace string
	}
	tests := []struct {
		name string
		args args
		want *serversTransports
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

			if got := newServersTransports(tt.args.c, tt.args.namespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newServersTransports() = %v, want %v", got, tt.want)
			}
		})
	}
}
