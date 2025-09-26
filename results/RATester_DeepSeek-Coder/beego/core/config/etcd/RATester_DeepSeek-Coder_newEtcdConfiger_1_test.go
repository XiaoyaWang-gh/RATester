package etcd

import (
	"fmt"
	"reflect"
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestNewEtcdConfiger_1(t *testing.T) {
	type args struct {
		client *clientv3.Client
		prefix string
	}
	tests := []struct {
		name string
		args args
		want *EtcdConfiger
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

			if got := newEtcdConfiger(tt.args.client, tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEtcdConfiger() = %v, want %v", got, tt.want)
			}
		})
	}
}
