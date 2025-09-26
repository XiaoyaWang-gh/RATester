package migration

import (
	"fmt"
	"testing"
)

func TestRegister_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	m := &Migration{}
	err := Register(name, m)
	if err != nil {
		t.Errorf("Register failed, err: %v", err)
	}
}
