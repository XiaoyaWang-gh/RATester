package auth

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewTokenAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	additionalAuthScopes := []v1.AuthScope{"scope1", "scope2"}
	token := "test_token"
	auth := NewTokenAuth(additionalAuthScopes, token)

	if auth == nil {
		t.Fatal("NewTokenAuth returned nil")
	}

	if len(auth.additionalAuthScopes) != len(additionalAuthScopes) {
		t.Fatalf("Expected %d additional auth scopes, got %d", len(additionalAuthScopes), len(auth.additionalAuthScopes))
	}

	for i, scope := range additionalAuthScopes {
		if auth.additionalAuthScopes[i] != scope {
			t.Fatalf("Expected additional auth scope %s, got %s", scope, auth.additionalAuthScopes[i])
		}
	}

	if auth.token != token {
		t.Fatalf("Expected token %s, got %s", token, auth.token)
	}
}
