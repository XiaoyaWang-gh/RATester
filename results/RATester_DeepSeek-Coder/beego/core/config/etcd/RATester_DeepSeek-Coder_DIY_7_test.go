package etcd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestDIY_7(t *testing.T) {
	type fields struct {
		prefix       string
		client       *clientv3.Client
		BaseConfiger config.BaseConfiger
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
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

			e := &EtcdConfiger{
				prefix:       tt.fields.prefix,
				client:       tt.fields.client,
				BaseConfiger: tt.fields.BaseConfiger,
			}
			got, err := e.DIY(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EtcdConfiger.DIY() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EtcdConfiger.DIY() = %v, want %v", got, tt.want)
			}
		})
	}
}
