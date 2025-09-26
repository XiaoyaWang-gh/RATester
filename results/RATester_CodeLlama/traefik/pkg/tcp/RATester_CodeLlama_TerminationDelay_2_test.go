package tcp

import (
	"fmt"
	"testing"
)

func TestTerminationDelay_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := tcpDialer{}
	if got := d.TerminationDelay(); got != 0 {
		t.Errorf("TerminationDelay() = %v, want %v", got, 0)
	}
}
