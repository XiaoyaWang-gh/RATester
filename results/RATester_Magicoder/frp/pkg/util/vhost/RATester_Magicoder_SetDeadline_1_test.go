package vhost

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestSetDeadline_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := readOnlyConn{
		reader: strings.NewReader("test data"),
	}

	testTime := time.Now()
	err := conn.SetDeadline(testTime)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
