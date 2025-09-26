package cache

import (
	"fmt"
	"testing"
)

func Testdel_1(t *testing.T) {
	type args struct {
		key string
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

			tt.m.del(tt.args.key)
		})
	}
}
