package metrics

import (
	"fmt"
	"testing"
)

func TestSetDynamicConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ps := &prometheusState{}
	dynamicConfig := &dynamicConfig{}
	ps.SetDynamicConfig(dynamicConfig)
}
