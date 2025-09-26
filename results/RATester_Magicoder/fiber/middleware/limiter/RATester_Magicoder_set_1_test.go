package limiter

import (
	"fmt"
	"testing"
	"time"
)

func Testset_1(t *testing.T) {
	type args struct {
		key string
		it  *item
		exp time.Duration
	}
	tests := []struct {
		name string
		m    *manager
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

			tt.m.set(tt.args.key, tt.args.it, tt.args.exp)
		})
	}
}
