package auth

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"golang.org/x/oauth2/clientcredentials"
)

func TestGenerateAccessToken_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	auth := &OidcAuthProvider{
		additionalAuthScopes: []v1.AuthScope{},
		tokenGenerator: &clientcredentials.Config{
			ClientID:     "test_client_id",
			ClientSecret: "test_client_secret",
			TokenURL:     "test_token_url",
		},
	}

	accessToken, err := auth.generateAccessToken()
	if err != nil {
		t.Errorf("generateAccessToken() returned an error: %v", err)
	}

	if accessToken == "" {
		t.Error("generateAccessToken() returned an empty access token")
	}
}
