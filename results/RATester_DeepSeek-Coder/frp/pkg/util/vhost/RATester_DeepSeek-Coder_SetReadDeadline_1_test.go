package vhost

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestSetReadDeadline_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := readOnlyConn{
		reader: strings.NewReader("test"),
	}

	testTime := time.Now()
	err := conn.SetReadDeadline(testTime)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
