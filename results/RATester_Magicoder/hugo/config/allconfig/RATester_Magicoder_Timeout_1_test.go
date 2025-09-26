package allconfig

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &Config{
		C: &ConfigCompiled{
			Timeout: 10 * time.Second,
		},
	}

	cl := ConfigLanguage{config: config}

	if got := cl.Timeout(); got != config.C.Timeout {
		t.Errorf("Timeout() = %v, want %v", got, config.C.Timeout)
	}
}
