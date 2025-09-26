package nathole

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestGetRangePorts_1(t *testing.T) {
	type args struct {
		addrs      []string
		difference int
		maxNumber  int
	}
	tests := []struct {
		name string
		args args
		want []msg.PortsRange
	}{
		{
			name: "test1",
			args: args{
				addrs:      []string{"127.0.0.1:10000", "127.0.0.1:10001", "127.0.0.1:10002"},
				difference: 100,
				maxNumber:  100,
			},
			want: []msg.PortsRange{
				{
					From: 9999,
					To:   10000,
				},
				{
					From: 10001,
					To:   10002,
				},
				{
					From: 10003,
					To:   10004,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getRangePorts(tt.args.addrs, tt.args.difference, tt.args.maxNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRangePorts() = %v, want %v", got, tt.want)
			}
		})
	}
}
