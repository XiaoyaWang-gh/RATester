package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMounts_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &moduleAdapter{
		mounts: []Mount{
			{
				Source: "scss",
			},
			{
				Source: "scss",
			},
		},
	}

	if got := m.Mounts(); !reflect.DeepEqual(got, m.mounts) {
		t.Errorf("Mounts() = %v, want %v", got, m.mounts)
	}
}
