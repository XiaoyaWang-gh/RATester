package gateway

import (
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestGetCertificateBlocks_2(t *testing.T) {
	type args struct {
		secret     *corev1.Secret
		namespace  string
		secretName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
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

			got, got1, err := getCertificateBlocks(tt.args.secret, tt.args.namespace, tt.args.secretName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCertificateBlocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCertificateBlocks() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getCertificateBlocks() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
