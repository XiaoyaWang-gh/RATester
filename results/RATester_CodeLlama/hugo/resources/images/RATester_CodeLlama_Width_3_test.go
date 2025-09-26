package images

import (
	"fmt"
	"testing"
)

func TestWidth_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := &Image{}
	i.initConfig()
	i.config.Width = 100
	if i.Width() != 100 {
		t.Errorf("Expected 100, got %d", i.Width())
	}
}
