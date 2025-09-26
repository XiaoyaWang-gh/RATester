package transport

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestNewMessageTransporter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sendCh := make(chan msg.Message)
	transporter := NewMessageTransporter(sendCh)
	if transporter == nil {
		t.Fatal("NewMessageTransporter failed")
	}
}
