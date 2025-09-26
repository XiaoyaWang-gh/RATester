package client

import (
	"context"
	"fmt"
	"net"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/zeebo/assert"
)

func TestNewControl_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sessionCtx := &SessionContext{
		Common: &v1.ClientCommonConfig{
			ServerAddr: "127.0.0.1:7000",
		},
		RunID: "test",
		Conn:  &net.TCPConn{},
	}

	ctl, err := NewControl(ctx, sessionCtx)
	assert.NoError(t, err)
	assert.NotNil(t, ctl)

	ctl.Close()
}
