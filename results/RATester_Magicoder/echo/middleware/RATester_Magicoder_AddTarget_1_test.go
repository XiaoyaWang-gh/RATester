package middleware

import (
	"fmt"
	"net/url"
	"sync"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestAddTarget_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &commonBalancer{
		targets: []*ProxyTarget{},
		mutex:   sync.Mutex{},
	}

	target1 := &ProxyTarget{
		Name: "target1",
		URL:  &url.URL{},
		Meta: echo.Map{},
	}

	target2 := &ProxyTarget{
		Name: "target2",
		URL:  &url.URL{},
		Meta: echo.Map{},
	}

	result1 := b.AddTarget(target1)
	if !result1 {
		t.Errorf("Expected true, got %v", result1)
	}

	result2 := b.AddTarget(target2)
	if !result2 {
		t.Errorf("Expected true, got %v", result2)
	}

	result3 := b.AddTarget(target1)
	if result3 {
		t.Errorf("Expected false, got %v", result3)
	}
}
