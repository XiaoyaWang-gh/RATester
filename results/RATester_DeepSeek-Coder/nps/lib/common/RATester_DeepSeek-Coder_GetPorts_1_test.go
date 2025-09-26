package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetPorts_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []int
	}{
		{"single port", "80", []int{80}},
		{"range of ports", "80-85", []int{80, 81, 82, 83, 84, 85}},
		{"mixed ports", "80,81-83,85", []int{80, 81, 82, 83, 85}},
		{"invalid ports", "80,81-83-85,90", []int{80, 81, 82, 83, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetPorts(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPorts() = %v, want %v", got, tt.want)
			}
		})
	}
}
