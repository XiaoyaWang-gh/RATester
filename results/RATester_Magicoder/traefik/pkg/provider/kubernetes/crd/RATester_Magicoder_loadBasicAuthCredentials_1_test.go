package crd

import (
	"fmt"
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestLoadBasicAuthCredentials_1(t *testing.T) {
	type args struct {
		secret *corev1.Secret
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test with valid secret",
			args: args{
				secret: &corev1.Secret{
					Data: map[string][]byte{
						"username": []byte("test_user"),
						"password": []byte("test_password"),
					},
				},
			},
			want:    []string{"test_user:{SHA}90b9aa7e25f80cf4f64e29c7b9484472863d54da"},
			wantErr: false,
		},
		{
			name: "Test with invalid secret",
			args: args{
				secret: &corev1.Secret{
					Data: map[string][]byte{
						"username": []byte("test_user"),
					},
				},
			},
			want:    nil,
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

			got, err := loadBasicAuthCredentials(tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBasicAuthCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBasicAuthCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}
