package crd

import (
	"fmt"
	"testing"
)

func TestLoadCASecret_1(t *testing.T) {
	type args struct {
		namespace  string
		secretName string
		k8sClient  Client
	}
	tests := []struct {
		name    string
		args    args
		want    string
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

			got, err := loadCASecret(tt.args.namespace, tt.args.secretName, tt.args.k8sClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadCASecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("loadCASecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
