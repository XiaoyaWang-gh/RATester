package client

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/lib/config"
)

func TestNewRPClient_1(t *testing.T) {
	type args struct {
		svraddr        string
		vKey           string
		bridgeConnType string
		proxyUrl       string
		cnf            *config.Config
		disconnectTime int
	}
	tests := []struct {
		name string
		args args
		want *TRPClient
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

			got := NewRPClient(tt.args.svraddr, tt.args.vKey, tt.args.bridgeConnType, tt.args.proxyUrl, tt.args.cnf, tt.args.disconnectTime)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRPClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
