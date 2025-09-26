package etcd

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Testreader_1(t *testing.T) {
	type fields struct {
		prefix       string
		client       *clientv3.Client
		BaseConfiger config.BaseConfiger
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
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

			e := &EtcdConfiger{
				prefix:       tt.fields.prefix,
				client:       tt.fields.client,
				BaseConfiger: tt.fields.BaseConfiger,
			}
			got, err := e.reader(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EtcdConfiger.reader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EtcdConfiger.reader() = %v, want %v", got, tt.want)
			}
		})
	}
}
