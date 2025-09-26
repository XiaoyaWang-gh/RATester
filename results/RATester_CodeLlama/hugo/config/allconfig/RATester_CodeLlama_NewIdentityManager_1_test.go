package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestNewIdentityManager_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	name := "test"
	m := c.NewIdentityManager(name)
	if m == identity.NopManager {
		t.Errorf("NewIdentityManager() = %v, want %v", m, identity.NopManager)
	}
}
