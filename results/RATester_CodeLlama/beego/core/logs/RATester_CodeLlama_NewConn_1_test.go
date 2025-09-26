package logs

import (
	"fmt"
	"testing"
)

func TestNewConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := NewConn()
	if conn == nil {
		t.Error("NewConn() failed")
	}
}
