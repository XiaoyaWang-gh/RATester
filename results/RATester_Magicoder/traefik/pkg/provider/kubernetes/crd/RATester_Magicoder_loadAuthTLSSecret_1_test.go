package crd

import (
	"fmt"
	"testing"
)

func TestLoadAuthTLSSecret_1(t *testing.T) {
	type args struct {
		namespace  string
		secretName string
		k8sClient  Client
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

			got, got1, err := loadAuthTLSSecret(tt.args.namespace, tt.args.secretName, tt.args.k8sClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadAuthTLSSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("loadAuthTLSSecret() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("loadAuthTLSSecret() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
