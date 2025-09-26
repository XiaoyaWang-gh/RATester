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
		{
			Name: "target3",
			URL:  &url.URL{Scheme: "http", Host: "localhost:8082"},
		},
	}
	b := NewRandomBalancer(targets)
	if b == nil {
		t.Fatal("NewRandomBalancer() = nil")
	}
	if len(b.(*randomBalancer).targets) != len(targets) {
		t.Errorf("NewRandomBalancer() = %v, want %v", len(b.(*randomBalancer).targets), len(targets))
	}
}
