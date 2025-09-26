package auth

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/pkg/util/util"
)

func TestSetNewWorkConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	auth := &TokenAuthSetterVerifier{
		additionalAuthScopes: []v1.AuthScope{v1.AuthScopeNewWorkConns},
		token:                "token",
	}
	newWorkConnMsg := &msg.NewWorkConn{}
	err := auth.SetNewWorkConn(newWorkConnMsg)
	if err != nil {
		t.Fatal(err)
	}
	if newWorkConnMsg.Timestamp == 0 {
		t.Fatal("newWorkConnMsg.Timestamp is 0")
	}
	if newWorkConnMsg.PrivilegeKey != util.GetAuthKey(auth.token, newWorkConnMsg.Timestamp) {
		t.Fatal("newWorkConnMsg.PrivilegeKey is not equal to util.GetAuthKey(auth.token, newWorkConnMsg.Timestamp)")
	}
}
