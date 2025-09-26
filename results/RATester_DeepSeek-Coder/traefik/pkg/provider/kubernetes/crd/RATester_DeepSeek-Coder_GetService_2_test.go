package crd

import (
	"fmt"
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestGetService_2(t *testing.T) {
	type args struct {
		namespace string
		name      string
	}
	tests := []struct {
		name    string
		c       *clientWrapper
		args    args
		want    *corev1.Service
		want1   bool
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

			got, got1, err := tt.c.GetService(tt.args.namespace, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("clientWrapper.GetService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clientWrapper.GetService() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("clientWrapper.GetService() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
