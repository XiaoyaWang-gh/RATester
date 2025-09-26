package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/pkg/util/util"
)

func TestVerifyLogin_1(t *testing.T) {
	auth := &TokenAuthSetterVerifier{
		token: "test_token",
	}

	testCases := []struct {
		name    string
		login   *msg.Login
		wantErr bool
	}{
		{
			name: "valid token",
			login: &msg.Login{
				PrivilegeKey: util.GetAuthKey(auth.token, time.Now().Unix()),
				Timestamp:    time.Now().Unix(),
			},
			wantErr: false,
		},
		{
			name: "invalid token",
			login: &msg.Login{
				PrivilegeKey: "invalid_token",
				Timestamp:    time.Now().Unix(),
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := auth.VerifyLogin(tc.login)
			if (err != nil) != tc.wantErr {
				t.Errorf("VerifyLogin() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
