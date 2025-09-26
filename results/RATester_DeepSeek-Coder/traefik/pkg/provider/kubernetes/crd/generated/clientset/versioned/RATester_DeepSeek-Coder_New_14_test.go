package versioned

import (
	"fmt"
	"reflect"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestNew_14(t *testing.T) {
	type args struct {
		c rest.Interface
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

			if got := New(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
