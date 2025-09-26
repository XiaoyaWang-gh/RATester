package fiber

import (
	"fmt"
	"testing"
)

func TestExecuteOnForkHooks_1(t *testing.T) {
	type args struct {
		pid int
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

			h := &Hooks{}
			h.executeOnForkHooks(tt.args.pid)
		})
	}
}
