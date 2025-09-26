package cssjs

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources"
)

func TestNewPostCSSClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rs := &resources.Spec{}
	client := NewPostCSSClient(rs)

	if client.rs != rs {
		t.Errorf("Expected client.rs to be %v, but got %v", rs, client.rs)
	}
}
