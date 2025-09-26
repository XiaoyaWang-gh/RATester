package orm

import (
	"fmt"
	"testing"
)

func TestRegisterDriver_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	driverName := "mysql"
	typ := DriverType(1)
	err := RegisterDriver(driverName, typ)
	if err != nil {
		t.Errorf("RegisterDriver error: %v", err)
	}
}
