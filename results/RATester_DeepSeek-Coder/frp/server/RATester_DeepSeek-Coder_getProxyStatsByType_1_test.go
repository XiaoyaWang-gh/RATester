package server

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetProxyStatsByType_1(t *testing.T) {
	type args struct {
		proxyType string
	}
	tests := []struct {
		name           string
		args           args
		wantProxyInfos []*ProxyStatsInfo
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

			svr := &Service{}
			if gotProxyInfos := svr.getProxyStatsByType(tt.args.proxyType); !reflect.DeepEqual(gotProxyInfos, tt.wantProxyInfos) {
				t.Errorf("getProxyStatsByType() = %v, want %v", gotProxyInfos, tt.wantProxyInfos)
			}
		})
	}
}
