package nathole

import (
	"fmt"
	"testing"
)

func TestDoSTUNRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	conn := &discoverConn{}
	addr := "127.0.0.1:3478"

	// when
	_, err := conn.doSTUNRequest(addr)

	// then
	if err != nil {
		t.Errorf("doSTUNRequest() error = %v", err)
		return
	}
}
