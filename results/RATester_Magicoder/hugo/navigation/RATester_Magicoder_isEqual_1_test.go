package navigation

import (
	"fmt"
	"testing"
)

func TestisEqual_1(t *testing.T) {
	type args struct {
		inme *MenuEntry
	}
	tests := []struct {
		name string
		m    *MenuEntry
		args args
		want bool
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

			if got := tt.m.isEqual(tt.args.inme); got != tt.want {
				t.Errorf("MenuEntry.isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
