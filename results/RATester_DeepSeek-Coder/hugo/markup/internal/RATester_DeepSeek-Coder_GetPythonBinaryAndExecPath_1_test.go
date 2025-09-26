package internal

import (
	"fmt"
	"testing"
)

func TestGetPythonBinaryAndExecPath_1(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := GetPythonBinaryAndExecPath()
			if got != tt.want {
				t.Errorf("GetPythonBinaryAndExecPath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetPythonBinaryAndExecPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
