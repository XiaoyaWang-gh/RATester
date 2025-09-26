package rate

import (
	"fmt"
	"testing"
)

func TestStop_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Rate{
		bucketSize: 10,
		stopChan:   make(chan bool),
	}
	s.Start()
	s.Stop()
	if s.stopChan != nil {
		t.Errorf("s.stopChan should be nil")
	}
}
