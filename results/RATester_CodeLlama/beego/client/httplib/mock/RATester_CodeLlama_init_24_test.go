package mock

import (
	"fmt"
	"testing"
)

func TestInit_24(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	InitMockSetting()
}
