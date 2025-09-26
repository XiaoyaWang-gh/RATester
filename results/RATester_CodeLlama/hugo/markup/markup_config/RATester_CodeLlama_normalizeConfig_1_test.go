package markup_config

import (
	"fmt"
	"testing"
)

func TestNormalizeConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TEST_CODE_HERE
}
