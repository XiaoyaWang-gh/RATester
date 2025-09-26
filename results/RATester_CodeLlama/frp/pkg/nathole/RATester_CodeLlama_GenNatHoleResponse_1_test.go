package nathole

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
	"github.com/zeebo/assert"
)

func TestGenNatHoleResponse_1(t *testing.T) {
	c := &Controller{
		clientCfgs: map[string]*ClientCfg{
			"client1": {
				name: "client1",
				sk:   "sk1",
			},
			"client2": {
				name: "client2",
				sk:   "sk2",
			},
		},
		sessions: map[string]*Session{
			"session1": {
				sid: "session1",
			},
			"session2": {
				sid: "session2",
			},
		},
	}
	testCases := []struct {
		name     string
		input    string
		expected *msg.NatHoleResp
	}{
		{
			name:     "normal",
			input:    "session1",
			expected: &msg.NatHoleResp{Sid: "session1"},
		},
		{
			name:     "nil session",
			input:    "session3",
			expected: &msg.NatHoleResp{Error: "session not found"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual := c.GenNatHoleResponse("transactionID", c.sessions[tc.input], "error")
			assert.Equal(t, tc.expected, actual)
		})
	}
}
