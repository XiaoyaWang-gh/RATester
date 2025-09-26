package ssh

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/client/proxy"
	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestCreateSuccessInfo_1(t *testing.T) {
	type args struct {
		user string
		pc   v1.ProxyConfigurer
		ps   *proxy.WorkingStatus
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				user: "test",
				pc:   &v1.ProxyBaseConfig{},
				ps: &proxy.WorkingStatus{
					Name:       "test",
					Type:       "test",
					Phase:      "test",
					Err:        "test",
					Cfg:        &v1.ProxyBaseConfig{},
					RemoteAddr: "test",
				},
			},
			want: "\nfrp (via SSH) (Ctrl+C to quit)\n\nUser: test\nProxyName: test\nType: test\nRemoteAddress: test\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := createSuccessInfo(tt.args.user, tt.args.pc, tt.args.ps); got != tt.want {
				t.Errorf("createSuccessInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
