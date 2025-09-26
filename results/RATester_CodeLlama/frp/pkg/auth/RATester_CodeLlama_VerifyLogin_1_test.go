package auth

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestVerifyLogin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	auth := &TokenAuthSetterVerifier{
		token: "token",
	}
	loginMsg := &msg.Login{
		PrivilegeKey: "token",
	}
	err := auth.VerifyLogin(loginMsg)
	if err != nil {
		t.Errorf("VerifyLogin error: %v", err)
	}

	loginMsg.PrivilegeKey = "wrong_token"
	err = auth.VerifyLogin(loginMsg)
	if err == nil {
		t.Errorf("VerifyLogin should return error when token is wrong")
	}
}
