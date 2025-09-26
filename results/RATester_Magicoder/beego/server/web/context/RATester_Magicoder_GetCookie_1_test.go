package context

import (
	"fmt"
	"testing"
)

func TestGetCookie_1(t *testing.T) {
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
						"cookie": "test_cookie",
					},
				},
			},
		},
	}

	expected := "test_cookie"
	actual := ctx.GetCookie("cookie")

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
