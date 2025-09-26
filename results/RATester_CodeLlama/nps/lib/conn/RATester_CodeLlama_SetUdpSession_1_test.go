package conn

import (
	"fmt"
	"testing"

	"github.com/xtaci/kcp-go"
)

func TestSetUdpSession_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sess := &kcp.UDPSession{}
	SetUdpSession(sess)
	// TODO: check if the session is set correctly
}
