package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/client/proxy"
)

func TestGetProxyStatus_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		s    *statusExporterImpl
		args args
		want *proxy.WorkingStatus
		ok   bool
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

			got, ok := tt.s.GetProxyStatus(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) || ok != tt.ok {
				t.Errorf("statusExporterImpl.GetProxyStatus() = %v, %v, want %v, %v", got, ok, tt.want, tt.ok)
			}
		})
	}
}
