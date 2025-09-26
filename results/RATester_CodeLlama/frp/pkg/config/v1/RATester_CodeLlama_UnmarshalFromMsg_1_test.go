package v1

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
	"gotest.tools/assert"
)

func TestUnmarshalFromMsg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &SUDPProxyConfig{}
	m := &msg.NewProxy{
		ProxyName: "test",
		ProxyType: "sudp",
		Sk:        "sk",
		AllowUsers: []string{
			"user1",
			"user2",
		},
	}
	c.UnmarshalFromMsg(m)
	assert.Equal(t, "test", c.Name)
	assert.Equal(t, "sudp", c.Type)
	assert.Equal(t, "sk", c.Secretkey)
	assert.Equal(t, []string{"user1", "user2"}, c.AllowUsers)
}
