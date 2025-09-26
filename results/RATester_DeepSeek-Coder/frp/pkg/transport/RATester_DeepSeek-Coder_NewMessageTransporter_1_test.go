package transport

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestNewMessageTransporter_1(t *testing.T) {
	type args struct {
		sendCh chan msg.Message
	}
	tests := []struct {
		name string
		args args
		want MessageTransporter
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

			if got := NewMessageTransporter(tt.args.sendCh); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessageTransporter() = %v, want %v", got, tt.want)
			}
		})
	}
}
