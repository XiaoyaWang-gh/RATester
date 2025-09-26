package env

import (
	"fmt"
	"testing"
)

func TestGetAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	envs := GetAll()
	if len(envs) != 0 {
		t.Errorf("GetAll() = %v, want %v", len(envs), 0)
	}
}
