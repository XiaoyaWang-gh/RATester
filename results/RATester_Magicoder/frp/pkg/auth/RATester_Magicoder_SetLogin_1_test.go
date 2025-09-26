package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/pkg/util/util"
)

func TestSetLogin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	auth := &TokenAuthSetterVerifier{
		token: "test_token",
	}
	loginMsg := &msg.Login{
		Timestamp: time.Now().Unix(),
	}
	err := auth.SetLogin(loginMsg)
	if err != nil {
		t.Errorf("SetLogin failed, err: %v", err)
	}
	if loginMsg.PrivilegeKey != util.GetAuthKey(auth.token, loginMsg.Timestamp) {
		t.Errorf("PrivilegeKey not set correctly, expected: %s, got: %s", util.GetAuthKey(auth.token, loginMsg.Timestamp), loginMsg.PrivilegeKey)
	}
}
