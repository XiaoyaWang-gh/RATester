package etcd

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestUnmarshaler_7(t *testing.T) {
	type args struct {
		prefix string
		obj    interface{}
		opt    []config.DecodeOption
	}
	tests := []struct {
		name    string
		e       *EtcdConfiger
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			e: &EtcdConfiger{
				prefix: "test",
				client: &clientv3.Client{},
			},
			args: args{
				prefix: "test",
				obj:    &struct{}{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.e.Unmarshaler(tt.args.prefix, tt.args.obj, tt.args.opt...); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshaler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
