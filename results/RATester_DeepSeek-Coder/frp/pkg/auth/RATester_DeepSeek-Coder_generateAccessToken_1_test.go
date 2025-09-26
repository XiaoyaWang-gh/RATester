package auth

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"golang.org/x/oauth2/clientcredentials"
)

func TestGenerateAccessToken_1(t *testing.T) {
	type fields struct {
		additionalAuthScopes []v1.AuthScope
		tokenGenerator       *clientcredentials.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
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

			auth := &OidcAuthProvider{
				additionalAuthScopes: tt.fields.additionalAuthScopes,
				tokenGenerator:       tt.fields.tokenGenerator,
			}
			got, err := auth.generateAccessToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("OidcAuthProvider.generateAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OidcAuthProvider.generateAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
