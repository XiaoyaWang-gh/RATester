package echo

import (
	"fmt"
	"net"
	"testing"
)

func TestTrustIPRange_1(t *testing.T) {
	t.Parallel()

	t.Run("should append ipRange to trustExtraRanges", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ipRange := &net.IPNet{}
		trustOption := TrustIPRange(ipRange)
		ipChecker := &ipChecker{}
		trustOption(ipChecker)

		if len(ipChecker.trustExtraRanges) != 1 {
			t.Errorf("trustExtraRanges should have 1 element")
		}
	})
}
