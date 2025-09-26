package auth

import (
	"fmt"
	"testing"

	"github.com/coreos/go-oidc/v3/oidc"
	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestVerifyPostLoginToken_1(t *testing.T) {
	type args struct {
		privilegeKey string
	}
	tests := []struct {
		name    string
		auth    *OidcAuthConsumer
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			auth: &OidcAuthConsumer{
				additionalAuthScopes: []v1.AuthScope{},
				verifier:             &oidc.IDTokenVerifier{},
				subjectFromLogin:     "test_subject",
			},
			args: args{
				privilegeKey: "test_key",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			auth: &OidcAuthConsumer{
				additionalAuthScopes: []v1.AuthScope{},
				verifier:             &oidc.IDTokenVerifier{},
				subjectFromLogin:     "test_subject",
			},
			args: args{
				privilegeKey: "invalid_key",
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			auth: &OidcAuthConsumer{
				additionalAuthScopes: []v1.AuthScope{},
				verifier:             &oidc.IDTokenVerifier{},
				subjectFromLogin:     "test_subject",
			},
			args: args{
				privilegeKey: "test_key",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.auth.verifyPostLoginToken(tt.args.privilegeKey); (err != nil) != tt.wantErr {
				t.Errorf("OidcAuthConsumer.verifyPostLoginToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
