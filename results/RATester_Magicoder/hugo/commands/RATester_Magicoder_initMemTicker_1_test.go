package commands

import (
	"fmt"
	"testing"
	"time"
)

func TestinitMemTicker_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hb := &hugoBuilder{}
	stop := hb.initMemTicker()
	time.Sleep(6 * time.Second)
	stop()
}
