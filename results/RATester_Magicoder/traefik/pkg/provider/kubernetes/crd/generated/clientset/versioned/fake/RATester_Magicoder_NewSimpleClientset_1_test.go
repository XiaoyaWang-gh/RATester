package fake

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
)

func TestNewSimpleClientset_1(t *testing.T) {
	type args struct {
		objects []runtime.Object
	}
	tests := []struct {
		name string
		args args
		want *Clientset
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

			if got := NewSimpleClientset(tt.args.objects...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimpleClientset() = %v, want %v", got, tt.want)
			}
		})
	}
}
