package conn

import (
	"fmt"
	"testing"
)

func TestNewLenConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := NewLenConn(nil)
	if conn == nil {
		t.Errorf("NewLenConn() = nil, want not nil")
	}
}
