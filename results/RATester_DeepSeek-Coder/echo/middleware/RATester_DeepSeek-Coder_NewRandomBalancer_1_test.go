package middleware

import (
	"fmt"
	"net/url"
	"testing"
)

func TestNewRandomBalancer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	targets := []*ProxyTarget{
		{
			Name: "target1",
			URL:  &url.URL{Scheme: "http", Host: "localhost:8080"},
		},
		{
			Name: "target2",
			URL:  &url.URL{Scheme: "http", Host: "localhost:8081"},
		},
	}

	b := NewRandomBalancer(targets)

	if b == nil {
		t.Errorf("Expected a balancer, got nil")
	}

	if len(b.(*randomBalancer).targets) != len(targets) {
		t.Errorf("Expected %d targets, got %d", len(targets), len(b.(*randomBalancer).targets))
	}

	for i, target := range targets {
		if b.(*randomBalancer).targets[i].Name != target.Name {
			t.Errorf("Expected target name %s, got %s", target.Name, b.(*randomBalancer).targets[i].Name)
		}
		if b.(*randomBalancer).targets[i].URL.String() != target.URL.String() {
			t.Errorf("Expected target URL %s, got %s", target.URL.String(), b.(*randomBalancer).targets[i].URL.String())
		}
	}
}
