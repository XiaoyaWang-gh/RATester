package versioned

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestNewForConfigAndClient_1(t *testing.T) {
	type args struct {
		c          *rest.Config
		httpClient *http.Client
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

			got, err := NewForConfigAndClient(tt.args.c, tt.args.httpClient)
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
