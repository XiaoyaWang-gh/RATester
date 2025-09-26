package nathole

import (
	"fmt"
	"testing"
)

func TestListLocalIPsForNatHole_1(t *testing.T) {
	t.Run("ListLocalIPsForNatHole", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ips, err := ListLocalIPsForNatHole(10)
		if err != nil {
			t.Errorf("ListLocalIPsForNatHole error: %v", err)
		}
		t.Logf("ips: %v", ips)
	})
}
