package integrity

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources"
)

func TestNew_49(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rs := &resources.Spec{}
	client := New(rs)

	if client.rs != rs {
		t.Errorf("Expected client.rs to be %v, but got %v", rs, client.rs)
	}
}
