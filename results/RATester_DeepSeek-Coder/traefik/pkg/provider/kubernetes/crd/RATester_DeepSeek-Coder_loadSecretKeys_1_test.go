package crd

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadSecretKeys_1(t *testing.T) {
	type args struct {
		k8sClient Client
		ns        string
		i         interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
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

			got, err := loadSecretKeys(tt.args.k8sClient, tt.args.ns, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadSecretKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadSecretKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
