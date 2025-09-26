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
						"username": []byte("testuser"),
						"password": []byte("testpass"),
					},
				},
			},
			want:    []string{"testuser:{SHA}9Gq5Tz9r7hxzJIUC8IzqIvIJFsI="},
			wantErr: false,
		},
		{
			name: "Test with missing username",
			args: args{
				secret: &corev1.Secret{
					Data: map[string][]byte{
						"password": []byte("testpass"),
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test with missing password",
			args: args{
				secret: &corev1.Secret{
					Data: map[string][]byte{
						"username": []byte("testuser"),
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
