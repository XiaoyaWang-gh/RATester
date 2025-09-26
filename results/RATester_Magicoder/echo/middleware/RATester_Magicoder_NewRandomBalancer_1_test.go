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
		{Name: "target1", URL: &url.URL{Scheme: "http", Host: "localhost:8080"}},
		{Name: "target2", URL: &url.URL{Scheme: "http", Host: "localhost:8081"}},
	}

	balancer := NewRandomBalancer(targets)

	if balancer == nil {
		t.Error("Expected balancer to be not nil")
	}

	target := balancer.Next(nil)
	if target == nil {
		t.Error("Expected target to be not nil")
	}

	if target.Name != "target1" && target.Name != "target2" {
		t.Errorf("Expected target name to be 'target1' or 'target2', got %s", target.Name)
	}

	if target.URL.String() != "http://localhost:8080" && target.URL.String() != "http://localhost:8081" {
		t.Errorf("Expected target URL to be 'http://localhost:8080' or 'http://localhost:8081', got %s", target.URL.String())
	}
}
