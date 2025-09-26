package auth

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestVerifyLogin_1(t *testing.T) {
	type args struct {
		m *msg.Login
	}
	tests := []struct {
		name    string
		auth    *TokenAuthSetterVerifier
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			auth := &TokenAuthSetterVerifier{
				additionalAuthScopes: tt.auth.additionalAuthScopes,
				token:                tt.auth.token,
			}
			if err := auth.VerifyLogin(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("TokenAuthSetterVerifier.VerifyLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
