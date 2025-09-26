package allconfig

import (
	"fmt"
	"testing"
)

func TestGetConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.GetConfig() != nil {
		t.Errorf("GetConfig() = %v, want nil", c.GetConfig())
	}
}
