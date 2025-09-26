package fiber

import (
	"fmt"
	"testing"
)

func TestSet_3(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name string
		c    *DefaultCtx
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

			tt.c.Set(tt.args.key, tt.args.val)
		})
	}
}
