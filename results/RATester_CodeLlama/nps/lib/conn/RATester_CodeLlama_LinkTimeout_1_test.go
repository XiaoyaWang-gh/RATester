package conn

import (
	"fmt"
	"testing"
	"time"
)

func TestLinkTimeout_1(t *testing.T) {
	t.Run("test case 1:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		opt := &Options{}
		LinkTimeout(time.Second)(opt)
		if opt.Timeout != time.Second {
			t.Errorf("expected %v, actual %v", time.Second, opt.Timeout)
		}
	})
}
