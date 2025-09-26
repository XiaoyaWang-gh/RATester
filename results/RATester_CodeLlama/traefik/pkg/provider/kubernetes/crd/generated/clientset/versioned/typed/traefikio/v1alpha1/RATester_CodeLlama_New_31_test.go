package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestNew_31(t *testing.T) {
	type args struct {
		c rest.Interface
	}
	tests := []struct {
		name string
		args args
		want *TraefikV1alpha1Client
	}{
		{
			name: "test",
			args: args{
				c: nil,
			},
			want: &TraefikV1alpha1Client{
				restClient: nil,
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

			if got := New(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
