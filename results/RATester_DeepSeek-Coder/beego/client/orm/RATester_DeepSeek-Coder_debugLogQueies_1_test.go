package orm

import (
	"fmt"
	"testing"
	"time"
)

func TestDebugLogQueies_1(t *testing.T) {
	type args struct {
		alias     *alias
		operation string
		query     string
		t         time.Time
		err       error
		args      []interface{}
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

			debugLogQueies(tt.args.alias, tt.args.operation, tt.args.query, tt.args.t, tt.args.err, tt.args.args...)
		})
	}
}
