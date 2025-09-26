package crd

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/client-go/rest"
)

func TestCreateClientFromConfig_3(t *testing.T) {
	type args struct {
		c *rest.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *clientWrapper
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

			got, err := createClientFromConfig(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("createClientFromConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createClientFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
