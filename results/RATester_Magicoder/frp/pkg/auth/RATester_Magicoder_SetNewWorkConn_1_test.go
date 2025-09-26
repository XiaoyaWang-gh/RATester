package auth

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
)

func TestSetNewWorkConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	auth := &TokenAuthSetterVerifier{
		additionalAuthScopes: []v1.AuthScope{v1.AuthScopeNewWorkConns},
		token:                "test_token",
	}

	newWorkConnMsg := &msg.NewWorkConn{
		RunID: "test_run_id",
	}

	err := auth.SetNewWorkConn(newWorkConnMsg)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if newWorkConnMsg.Timestamp == 0 {
		t.Errorf("Timestamp not set")
	}

	if newWorkConnMsg.PrivilegeKey == "" {
		t.Errorf("PrivilegeKey not set")
	}
}
