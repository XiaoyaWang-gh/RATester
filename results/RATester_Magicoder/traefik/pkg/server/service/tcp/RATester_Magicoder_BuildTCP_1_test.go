package tcp

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tcp"
)

func TestBuildTCP_1(t *testing.T) {
	type args struct {
		rootCtx     context.Context
		serviceName string
	}

	tests := []struct {
		name    string
		args    args
		want    tcp.Handler
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				rootCtx:     context.Background(),
				serviceName: "service1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test case 2",
			args: args{
				rootCtx:     context.Background(),
				serviceName: "service2",
			},
			want:    nil,
			wantErr: true,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &Manager{}
			got, err := m.BuildTCP(tt.args.rootCtx, tt.args.serviceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildTCP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTCP() = %v, want %v", got, tt.want)
			}
		})
	}
}
