package visitor

import (
	"fmt"
	"testing"
)

func TestClose_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := &QUICTunnelSession{}
	qs.Close()
}
