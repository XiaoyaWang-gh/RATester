package mock

import (
	"fmt"
	"testing"
)

func TestInitMockSetting_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	InitMockSetting()

	// Add assertions here to verify the correctness of the method
}
