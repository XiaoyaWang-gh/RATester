package gateway

import (
	"fmt"
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestGetService_1(t *testing.T) {
	type args struct {
		namespace string
		name      string
	}
	tests := []struct {
		name    string
		args    args
		want    *corev1.Service
		want1   bool
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				namespace: "test-namespace",
				name:      "test-service",
			},
			want:    &corev1.Service{},
			want1:   true,
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				namespace: "non-watched-namespace",
				name:      "test-service",
			},
			want:    nil,
			want1:   false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &clientWrapper{}
			got, got1, err := c.GetService(tt.args.namespace, tt.args.name)
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
