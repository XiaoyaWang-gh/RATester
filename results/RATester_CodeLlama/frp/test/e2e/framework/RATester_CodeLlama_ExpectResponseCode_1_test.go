package framework

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestExpectResponseCode_1(t *testing.T) {
	t.Run("ExpectResponseCode", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		resp := &request.Response{
			Code: 200,
		}
		if !ExpectResponseCode(200)(resp) {
			t.Errorf("ExpectResponseCode(200) failed")
		}
	})
}
