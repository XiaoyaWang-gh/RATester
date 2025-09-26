package tcp

import (
	"fmt"
	"testing"
	"time"
)

func TestTerminationDelay_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := tcpDialer{
		terminationDelay: 10 * time.Second,
	}

	expected := 10 * time.Second
	actual := d.TerminationDelay()

	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
