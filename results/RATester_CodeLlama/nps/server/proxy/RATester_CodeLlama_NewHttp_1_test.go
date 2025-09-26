package proxy

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"ehang.io/nps/bridge"
	"ehang.io/nps/lib/file"
)

func TestNewHttp_1(t *testing.T) {
	type args struct {
		bridge    *bridge.Bridge
		c         *file.Tunnel
		httpPort  int
		httpsPort int
		useCache  bool
		cacheLen  int
		addOrigin bool
	}
	tests := []struct {
		name string
		args args
		want *httpServer
	}{
		{
			name: "test",
			args: args{
				bridge:    &bridge.Bridge{},
				c:         &file.Tunnel{},
				httpPort:  0,
				httpsPort: 0,
				useCache:  false,
				cacheLen:  0,
				addOrigin: false,
			},
			want: &httpServer{
				BaseServer: BaseServer{
					task:   &file.Tunnel{},
					bridge: &bridge.Bridge{},
					Mutex:  sync.Mutex{},
				},
				httpPort:  0,
				httpsPort: 0,
				useCache:  false,
				cacheLen:  0,
				addOrigin: false,
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

			if got := NewHttp(tt.args.bridge, tt.args.c, tt.args.httpPort, tt.args.httpsPort, tt.args.useCache, tt.args.cacheLen, tt.args.addOrigin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttp() = %v, want %v", got, tt.want)
			}
		})
	}
}
