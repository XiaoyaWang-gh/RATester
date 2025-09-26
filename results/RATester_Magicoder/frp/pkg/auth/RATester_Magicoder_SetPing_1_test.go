package auth

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
)

func TestSetPing_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	auth := &TokenAuthSetterVerifier{
		additionalAuthScopes: []v1.AuthScope{v1.AuthScopeHeartBeats},
		token:                "test_token",
	}

	pingMsg := &msg.Ping{}
	err := auth.SetPing(pingMsg)
	if err != nil {
		t.Fatalf("SetPing failed: %v", err)
	}

	if pingMsg.Timestamp == 0 {
		t.Fatalf("Timestamp is not set")
	}

	if pingMsg.PrivilegeKey == "" {
		t.Fatalf("PrivilegeKey is not set")
	}
}
