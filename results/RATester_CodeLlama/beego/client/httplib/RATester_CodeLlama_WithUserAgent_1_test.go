package httplib

import (
	"fmt"
	"testing"
)

func TestWithUserAgent_1(t *testing.T) {
	t.Run("WithUserAgent", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		userAgent := "test"
		client := &Client{}
		WithUserAgent(userAgent)(client)
		if client.Setting.UserAgent != userAgent {
			t.Errorf("WithUserAgent failed")
		}
	})
}
