package hugolib

import (
	"fmt"
	"testing"
)

func Testerrorf_7(t *testing.T) {
	type args struct {
		err    error
		format string
		a      []any
	}
	tests := []struct {
		name    string
		p       *pageState
		args    args
		wantErr bool
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

			if err := tt.p.errorf(tt.args.err, tt.args.format, tt.args.a...); (err != nil) != tt.wantErr {
				t.Errorf("pageState.errorf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
