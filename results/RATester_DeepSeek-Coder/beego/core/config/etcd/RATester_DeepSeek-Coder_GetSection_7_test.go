package etcd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestGetSection_7(t *testing.T) {
	type fields struct {
		prefix       string
		client       *clientv3.Client
		BaseConfiger config.BaseConfiger
	}
	type args struct {
		section string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]string
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
			got, err := e.GetSection(tt.args.section)
			if (err != nil) != tt.wantErr {
				t.Errorf("EtcdConfiger.GetSection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EtcdConfiger.GetSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
