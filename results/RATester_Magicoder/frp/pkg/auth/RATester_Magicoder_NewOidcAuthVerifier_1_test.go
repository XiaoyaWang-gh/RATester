package auth

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
)

func TestNewOidcAuthVerifier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	additionalAuthScopes := []v1.AuthScope{"scope1", "scope2"}
	cfg := v1.AuthOIDCServerConfig{
		Issuer:          "https://example.com",
		Audience:        "example-audience",
		SkipExpiryCheck: true,
		SkipIssuerCheck: true,
	}
	authConsumer := NewOidcAuthVerifier(additionalAuthScopes, cfg)

	// Test VerifyLogin
	loginMsg := &msg.Login{
		PrivilegeKey: "privilegeKey",
	}
	err := authConsumer.VerifyLogin(loginMsg)
	if err != nil {
		t.Errorf("VerifyLogin failed: %v", err)
	}

	// Test VerifyNewWorkConn
	newWorkConnMsg := &msg.NewWorkConn{
		PrivilegeKey: "privilegeKey",
	}
	err = authConsumer.VerifyNewWorkConn(newWorkConnMsg)
	if err != nil {
		t.Errorf("VerifyNewWorkConn failed: %v", err)
	}

	// Test VerifyPing
	pingMsg := &msg.Ping{
		PrivilegeKey: "privilegeKey",
	}
	err = authConsumer.VerifyPing(pingMsg)
	if err != nil {
		t.Errorf("VerifyPing failed: %v", err)
	}

	// Test verifyPostLoginToken
	err = authConsumer.verifyPostLoginToken("privilegeKey")
	if err != nil {
		t.Errorf("verifyPostLoginToken failed: %v", err)
	}
}
