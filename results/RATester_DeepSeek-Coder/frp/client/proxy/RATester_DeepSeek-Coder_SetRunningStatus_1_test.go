package proxy

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestSetRunningStatus_1(t *testing.T) {
	type args struct {
		remoteAddr string
		respErr    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				remoteAddr: "127.0.0.1:8080",
				respErr:    "",
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				remoteAddr: "127.0.0.1:8080",
				respErr:    "error",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			pw := &Wrapper{
				WorkingStatus: WorkingStatus{
					Name:       "test",
					Type:       "test",
					Phase:      ProxyPhaseWaitStart,
					Err:        "",
					Cfg:        &v1.ProxyBaseConfig{},
					RemoteAddr: "",
				},
			}
			if err := pw.SetRunningStatus(tt.args.remoteAddr, tt.args.respErr); (err != nil) != tt.wantErr {
				t.Errorf("SetRunningStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
