package context

import (
	"fmt"
	"testing"
)

func TestGetSecureCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Input: &BeegoInput{
			Context: &Context{
				Input: &BeegoInput{
					data: map[interface{}]interface{}{
						"cookie": "value|timestamp|signature",
					},
				},
			},
		},
	}

	val, ok := ctx.GetSecureCookie("secret", "cookie")
	if !ok {
		t.Error("Expected cookie to be valid")
	}
	if val != "value" {
		t.Errorf("Expected value to be 'value', got %s", val)
	}
}
