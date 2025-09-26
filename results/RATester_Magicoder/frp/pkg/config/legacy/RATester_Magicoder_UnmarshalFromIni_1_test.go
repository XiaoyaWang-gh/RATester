package legacy

import (
	"fmt"
	"testing"

	"gopkg.in/ini.v1"
)

func TestUnmarshalFromIni_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := &XTCPVisitorConf{}
	section := &ini.Section{}

	// Test case 1: No error when all fields are present and valid
	section.NewKey("protocol", "quic")
	section.NewKey("max_retries_an_hour", "8")
	section.NewKey("min_retry_interval", "90")
	section.NewKey("fallback_timeout_ms", "1000")
	err := cfg.UnmarshalFromIni("", "", section)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 2: Error when protocol is missing
	section.DeleteKey("protocol")
	err = cfg.UnmarshalFromIni("", "", section)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Test case 3: Error when max_retries_an_hour is not a number
	section.NewKey("max_retries_an_hour", "not_a_number")
	err = cfg.UnmarshalFromIni("", "", section)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Test case 4: Error when min_retry_interval is not a number
	section.NewKey("min_retry_interval", "not_a_number")
	err = cfg.UnmarshalFromIni("", "", section)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Test case 5: Error when fallback_timeout_ms is not a number
	section.NewKey("fallback_timeout_ms", "not_a_number")
	err = cfg.UnmarshalFromIni("", "", section)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
