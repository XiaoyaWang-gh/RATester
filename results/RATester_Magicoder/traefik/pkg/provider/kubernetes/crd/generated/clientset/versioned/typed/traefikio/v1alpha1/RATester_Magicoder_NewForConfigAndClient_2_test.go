package v1alpha1

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestNewForConfigAndClient_2(t *testing.T) {
	type args struct {
		c *rest.Config
		h *http.Client
	}
	tests := []struct {
		name    string
		args    args
		want    *TraefikV1alpha1Client
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

			got, err := NewForConfigAndClient(tt.args.c, tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewForConfigAndClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewForConfigAndClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
