package session

import (
	"fmt"
	"testing"
)

func TestDelSession_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Session{}
	s.delSession()
}
