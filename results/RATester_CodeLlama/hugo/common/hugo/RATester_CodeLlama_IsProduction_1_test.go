package hugo

import (
	"fmt"
	"testing"
)

func TestIsProduction_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := HugoInfo{
		Environment: EnvironmentProduction,
	}
	if !i.IsProduction() {
		t.Errorf("Expected IsProduction to be true")
	}
}
