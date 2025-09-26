package legacy

import (
	"fmt"
	"testing"
)

func TestGetDefaultTokenConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := TokenConfig{
		Token: "",
	}
	result := getDefaultTokenConf()
	if result.Token != expected.Token {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
