package cssjs

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources"
)

func TestNewTailwindCSSClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rs := &resources.Spec{}
	client := NewTailwindCSSClient(rs)

	if client == nil {
		t.Error("Expected a TailwindCSSClient, but got nil")
	}

	if client.rs != rs {
		t.Errorf("Expected client.rs to be %v, but got %v", rs, client.rs)
	}
}
