package etcd

import (
	"fmt"
	"reflect"
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func Testget_29(t *testing.T) {
	type args struct {
		client *clientv3.Client
		key    string
	}
	tests := []struct {
		name    string
		args    args
		want    *clientv3.GetResponse
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

			got, err := get(tt.args.client, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}
