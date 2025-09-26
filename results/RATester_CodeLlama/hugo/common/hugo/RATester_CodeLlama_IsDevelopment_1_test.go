package hugo

import (
	"fmt"
	"testing"
)

func TestIsDevelopment_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := HugoInfo{
		Environment: EnvironmentDevelopment,
	}
	if !i.IsDevelopment() {
		t.Errorf("Expected IsDevelopment to be true")
	}
}
