package auth

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/pkg/util/util"
)

func TestSetPing_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	auth := &TokenAuthSetterVerifier{
		additionalAuthScopes: []v1.AuthScope{v1.AuthScopeHeartBeats},
		token:                "token",
	}
	pingMsg := &msg.Ping{}
	err := auth.SetPing(pingMsg)
	if err != nil {
		t.Fatal(err)
	}
	if pingMsg.Timestamp == 0 {
		t.Fatal("pingMsg.Timestamp is 0")
	}
	if pingMsg.PrivilegeKey != util.GetAuthKey(auth.token, pingMsg.Timestamp) {
		t.Fatal("pingMsg.PrivilegeKey is not equal to util.GetAuthKey(auth.token, pingMsg.Timestamp)")
	}
}
