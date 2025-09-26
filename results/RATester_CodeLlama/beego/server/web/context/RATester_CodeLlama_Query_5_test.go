package context

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestQuery_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Form: url.Values{
					"key": []string{"value"},
				},
			},
		},
	}
	if input.Query("key") != "value" {
		t.Errorf("Query() = %v, want %v", input.Query("key"), "value")
	}
}
