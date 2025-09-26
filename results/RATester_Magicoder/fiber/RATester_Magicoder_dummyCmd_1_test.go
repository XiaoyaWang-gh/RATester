package fiber

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func TestdummyCmd_1(t *testing.T) {
	tests := []struct {
		name string
		want *exec.Cmd
	}{
		{
			name: "Test case 1",
			want: &exec.Cmd{
				Path: "go",
				Args: []string{"go", "version"},
			},
		},
		{
			name: "Test case 2",
			want: &exec.Cmd{
				Path: "cmd",
				Args: []string{"cmd", "/C", "go", "version"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := dummyCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dummyCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
