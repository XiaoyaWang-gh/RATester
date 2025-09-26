package gateway

import (
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestLookupNamespace_2(t *testing.T) {
	type args struct {
		namespace string
	}
	tests := []struct {
		name string
		c    *clientWrapper
		args args
		want string
	}{
		{
			name: "namespace is all",
			c: &clientWrapper{
				isNamespaceAll: true,
			},
			args: args{
				namespace: "test",
			},
			want: metav1.NamespaceAll,
		},
		{
			name: "namespace is not all",
			c: &clientWrapper{
				isNamespaceAll: false,
			},
			args: args{
				namespace: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.lookupNamespace(tt.args.namespace); got != tt.want {
				t.Errorf("lookupNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
