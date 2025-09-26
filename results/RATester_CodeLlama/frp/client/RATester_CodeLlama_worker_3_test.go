package client

import (
	"fmt"
	"testing"
)

func TestWorker_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{}
	ctl.worker()
}
