package client

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/pkg/util/xlog"
)

func TestHandlePong_1(t *testing.T) {
	ctl := &Control{
		xl:       &xlog.Logger{},
		lastPong: atomic.Value{},
	}

	testCases := []struct {
		name     string
		inMsg    *msg.Pong
		expected string
	}{
		{
			name: "Error",
			inMsg: &msg.Pong{
				Error: "test error",
			},
			expected: "test error",
		},
		{
			name: "NoError",
			inMsg: &msg.Pong{
				Error: "",
			},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctl.handlePong(tc.inMsg)
			lastPong := ctl.lastPong.Load().(time.Time)
			if tc.expected != "" {
				ctl.xl.Errorf("Pong message contains error: %s", tc.expected)
				ctl.closeSession()
			}
			if lastPong.IsZero() {
				ctl.xl.Debugf("receive heartbeat from server")
			}
		})
	}
}
