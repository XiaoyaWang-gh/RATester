package net

import (
	"context"
	"fmt"
	"net"
	"testing"

	"gotest.tools/assert"
)

func TestNewContextConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	c := &net.TCPConn{}
	cc := NewContextConn(ctx, c)
	assert.Equal(t, c, cc.Conn)
	assert.Equal(t, ctx, cc.ctx)
}
