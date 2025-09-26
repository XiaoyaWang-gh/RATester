package logger

import (
	"fmt"
	"testing"
)

func TestBeforeHandlerFunc_1(t *testing.T) {
	type args struct {
		cfg Config
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

			beforeHandlerFunc(tt.args.cfg)
		})
	}
}
