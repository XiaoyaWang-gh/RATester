package tailscale

import (
	"fmt"
	"testing"
)

func TestThrottleDuration_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	if got := p.ThrottleDuration(); got != 0 {
		t.Errorf("ThrottleDuration() = %v, want %v", got, 0)
	}
}
