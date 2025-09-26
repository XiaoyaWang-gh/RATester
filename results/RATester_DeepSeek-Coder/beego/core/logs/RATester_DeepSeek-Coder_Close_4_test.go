package logs

import (
	"fmt"
	"testing"
)

func TestClose_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		asynchronous: true,
		closeChan:    make(chan struct{}),
	}

	go func() {
		<-bl.closeChan
		bl.wg.Done()
	}()

	bl.wg.Add(1)
	bl.Close()

	if _, ok := <-bl.closeChan; ok {
		t.Errorf("closeChan should be closed")
	}
}
