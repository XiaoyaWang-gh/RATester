package middleware

import (
	"fmt"
	"net/url"
	"testing"
)

func TestNewRoundRobinBalancer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	targets := []*ProxyTarget{
		{Name: "target1", URL: &url.URL{Scheme: "http", Host: "localhost:8080"}},
		{Name: "target2", URL: &url.URL{Scheme: "http", Host: "localhost:8081"}},
	}

	balancer := NewRoundRobinBalancer(targets)

	if balancer == nil {
		t.Error("Expected balancer to be not nil")
	}

	target := balancer.Next(nil)
	if target == nil {
		t.Error("Expected target to be not nil")
	}

	if target.Name != "target1" {
		t.Errorf("Expected target name to be 'target1', got %s", target.Name)
	}

	target = balancer.Next(nil)
	if target == nil {
		t.Error("Expected target to be not nil")
	}

	if target.Name != "target2" {
		t.Errorf("Expected target name to be 'target2', got %s", target.Name)
	}

	target = balancer.Next(nil)
	if target == nil {
		t.Error("Expected target to be not nil")
	}

	if target.Name != "target1" {
		t.Errorf("Expected target name to be 'target1', got %s", target.Name)
	}
}
