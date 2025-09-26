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
		{
			name: "single port",
			arg:  "80",
			want: []int{80},
		},
		{
			name: "range of ports",
			arg:  "80-85",
			want: []int{80, 81, 82, 83, 84, 85},
		},
		{
			name: "mixed ports",
			arg:  "80,85,443-445",
			want: []int{80, 85, 443, 444, 445},
		},
		{
			name: "invalid ports",
			arg:  "80,invalid,443-invalid",
			want: []int{80, 443},
		},
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
