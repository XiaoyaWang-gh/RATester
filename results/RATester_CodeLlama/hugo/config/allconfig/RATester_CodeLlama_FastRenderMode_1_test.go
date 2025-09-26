package allconfig

import (
	"fmt"
	"testing"
)

func TestFastRenderMode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{
		config: &Config{
			Internal: InternalConfig{
				FastRenderMode: true,
			},
		},
	}

	if !c.FastRenderMode() {
		t.Errorf("FastRenderMode should be true")
	}
}
