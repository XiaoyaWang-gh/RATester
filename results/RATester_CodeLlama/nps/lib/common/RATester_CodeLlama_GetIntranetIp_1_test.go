package common

import (
	"fmt"
	"testing"
)

func TestGetIntranetIp_1(t *testing.T) {
	t.Run("get intranet ip success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err, ip := GetIntranetIp()
		if err != nil {
			t.Errorf("get intranet ip error: %v", err)
		}
		t.Logf("get intranet ip success: %s", ip)
	})

	t.Run("get intranet ip error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err, ip := GetIntranetIp()
		if err == nil {
			t.Errorf("get intranet ip error: %v", err)
		}
		t.Logf("get intranet ip error: %s", ip)
	})
}
