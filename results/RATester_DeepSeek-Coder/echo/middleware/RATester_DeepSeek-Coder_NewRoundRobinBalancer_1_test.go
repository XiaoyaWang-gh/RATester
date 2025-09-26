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
		{
			Name: "target1",
			URL:  &url.URL{Scheme: "http", Host: "localhost:8080"},
		},
		{
			Name: "target2",
			URL:  &url.URL{Scheme: "http", Host: "localhost:8081"},
		},
	}

	balancer := NewRoundRobinBalancer(targets)

	if balancer == nil {
		t.Errorf("Expected balancer to be not nil")
	}

	nextTarget := balancer.Next(nil)
	if nextTarget == nil {
		t.Errorf("Expected nextTarget to be not nil")
	}

	if nextTarget.Name != "target1" {
		t.Errorf("Expected nextTarget.Name to be 'target1', got %s", nextTarget.Name)
	}

	nextTarget = balancer.Next(nil)
	if nextTarget == nil {
		t.Errorf("Expected nextTarget to be not nil")
	}

	if nextTarget.Name != "target2" {
		t.Errorf("Expected nextTarget.Name to be 'target2', got %s", nextTarget.Name)
	}

	nextTarget = balancer.Next(nil)
	if nextTarget == nil {
		t.Errorf("Expected nextTarget to be not nil")
	}

	if nextTarget.Name != "target1" {
		t.Errorf("Expected nextTarget.Name to be 'target1', got %s", nextTarget.Name)
	}
}
