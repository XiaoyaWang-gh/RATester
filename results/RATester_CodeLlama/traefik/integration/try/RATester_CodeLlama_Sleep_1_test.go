package try

import (
	"fmt"
	"testing"
	"time"
)

func TestSleep_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := time.Duration(1000)
	Sleep(d)
}
