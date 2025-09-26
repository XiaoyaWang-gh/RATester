package testenv

import (
	"fmt"
	"testing"
)

func TestGoToolPath_1(t *testing.T) {
	type args struct {
		t testing.TB
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := GoToolPath(tt.args.t); got != tt.want {
				t.Errorf("GoToolPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
