package toml

import (
	"fmt"
	"testing"
)

func TestDefaultInt_6(t *testing.T) {
	type args struct {
		key        string
		defaultVal int
	}
	tests := []struct {
		name string
		c    *configContainer
		args args
		want int
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

			if got := tt.c.DefaultInt(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("configContainer.DefaultInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
