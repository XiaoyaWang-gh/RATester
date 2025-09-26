package versioned

import (
	"fmt"
	"reflect"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestNewForConfig_1(t *testing.T) {
	type args struct {
		c *rest.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *Clientset
		wantErr bool
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

			got, err := NewForConfig(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewForConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewForConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
