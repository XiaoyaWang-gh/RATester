package externalversions

import (
	"fmt"
	reflect "reflect"
	"testing"
)

func TestWithNamespace_1(t *testing.T) {
	type args struct {
		namespace string
	}
	tests := []struct {
		name string
		args args
		want *sharedInformerFactory
	}{
		{
			name: "Test WithNamespace",
			args: args{
				namespace: "test-namespace",
			},
			want: &sharedInformerFactory{
				namespace: "test-namespace",
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

			if got := WithNamespace(tt.args.namespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
