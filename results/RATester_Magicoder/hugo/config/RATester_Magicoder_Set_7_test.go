package config

import (
	"fmt"
	"testing"
)

func TestSet_7(t *testing.T) {
	type args struct {
		k string
		v any
	}
	tests := []struct {
		name string
		args args
		want any
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

			c := &defaultConfigProvider{}
			c.Set(tt.args.k, tt.args.v)
		})
	}
}
