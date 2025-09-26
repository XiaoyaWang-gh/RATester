package htesting

import (
	"fmt"
	"testing"
	"time"
)

func TestBailOut_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			t.Log("recovered from panic:", r)
		}
	}()
	BailOut(time.Second)
	t.Log("BailOut did not panic")
}
