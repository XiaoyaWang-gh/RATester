package logs

import (
	"fmt"
	"testing"
)

func Testconnect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &connWriter{
		Net:  "tcp",
		Addr: "localhost:8080",
	}

	err := conn.connect()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if conn.innerWriter == nil {
		t.Error("Expected innerWriter to be not nil")
	}
}
