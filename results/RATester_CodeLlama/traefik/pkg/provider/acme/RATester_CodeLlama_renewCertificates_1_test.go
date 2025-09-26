package acme

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRenewCertificates_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	p := &Provider{}
	ctx := context.Background()
	renewPeriod := time.Duration(1)

	// when
	p.renewCertificates(ctx, renewPeriod)

	// then
	// TODO: add assertions
}
