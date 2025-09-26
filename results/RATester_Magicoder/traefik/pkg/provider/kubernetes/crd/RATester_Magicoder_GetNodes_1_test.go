package crd

import (
	"fmt"
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestGetNodes_1(t *testing.T) {
	tests := []struct {
		name    string
		want    []*corev1.Node
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

			c := &clientWrapper{}
			got, got1, err := c.GetNodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("clientWrapper.GetNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clientWrapper.GetNodes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("clientWrapper.GetNodes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
