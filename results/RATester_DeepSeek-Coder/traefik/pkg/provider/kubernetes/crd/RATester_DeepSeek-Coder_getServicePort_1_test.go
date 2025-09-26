package crd

import (
	"fmt"
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func TestGetServicePort_1(t *testing.T) {
	type args struct {
		svc  *corev1.Service
		port intstr.IntOrString
	}
	tests := []struct {
		name    string
		args    args
		want    *corev1.ServicePort
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

			got, err := getServicePort(tt.args.svc, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("getServicePort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getServicePort() = %v, want %v", got, tt.want)
			}
		})
	}
}
