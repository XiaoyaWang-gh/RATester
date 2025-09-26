package externalversions

import (
	"fmt"
	reflect "reflect"
	"testing"
	time "time"

	versioned "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/generated/clientset/versioned"
)

func TestNewSharedInformerFactory_1(t *testing.T) {
	type args struct {
		client        versioned.Interface
		defaultResync time.Duration
	}
	tests := []struct {
		name string
		args args
		want SharedInformerFactory
	}{
		{
			name: "test",
			args: args{
				client:        nil,
				defaultResync: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewSharedInformerFactory(tt.args.client, tt.args.defaultResync); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSharedInformerFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
