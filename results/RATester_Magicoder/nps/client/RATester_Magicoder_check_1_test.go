package client

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/file"
)

func Testcheck_1(t *testing.T) {
	type args struct {
		h *file.Health
	}
	tests := []struct {
		name string
		args args
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

			check(tt.args.h)
		})
	}
}
