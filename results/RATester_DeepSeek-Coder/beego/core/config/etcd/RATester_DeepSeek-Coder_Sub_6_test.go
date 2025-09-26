package etcd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestSub_6(t *testing.T) {
	type fields struct {
		prefix string
		client *clientv3.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    config.Configer
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

			e := &EtcdConfiger{
				prefix: tt.fields.prefix,
				client: tt.fields.client,
			}
			got, err := e.Sub(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EtcdConfiger.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EtcdConfiger.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
