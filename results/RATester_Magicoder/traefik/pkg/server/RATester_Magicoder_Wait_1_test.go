package server

import (
	"fmt"
	"testing"
	"time"
)

func TestWait_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		stopChan: make(chan bool),
	}

	go func() {
		time.Sleep(1 * time.Second)
		server.stopChan <- true
	}()

	server.Wait()
}
