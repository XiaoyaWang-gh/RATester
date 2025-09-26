package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/client-go/tools/cache"
)

func TestNewIngressRouteTCPLister_1(t *testing.T) {
	type args struct {
		indexer cache.Indexer
	}
	tests := []struct {
		name string
		args args
		want IngressRouteTCPLister
	}{
		{
			name: "test",
			args: args{
				indexer: nil,
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

			if got := NewIngressRouteTCPLister(tt.args.indexer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIngressRouteTCPLister() = %v, want %v", got, tt.want)
			}
		})
	}
}
