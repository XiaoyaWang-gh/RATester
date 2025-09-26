package maps

import (
	"fmt"
	"testing"
)

func TestSetInMap_1(t *testing.T) {
	type args struct {
		key    string
		mapKey string
		value  any
	}
	tests := []struct {
		name string
		c    *Scratch
		args args
		want string
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

			if got := tt.c.SetInMap(tt.args.key, tt.args.mapKey, tt.args.value); got != tt.want {
				t.Errorf("Scratch.SetInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
