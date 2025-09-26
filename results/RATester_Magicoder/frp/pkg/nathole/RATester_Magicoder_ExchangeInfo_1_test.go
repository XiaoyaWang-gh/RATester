package nathole

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/pkg/transport"
)

func TestExchangeInfo_1(t *testing.T) {
	type args struct {
		ctx         context.Context
		transporter transport.MessageTransporter
		laneKey     string
		m           msg.Message
		timeout     time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    *msg.NatHoleResp
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

			got, err := ExchangeInfo(tt.args.ctx, tt.args.transporter, tt.args.laneKey, tt.args.m, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExchangeInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExchangeInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
