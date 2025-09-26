package metric

import (
	"fmt"
	"testing"
	"time"
)

func TestRotate_1(t *testing.T) {
	type args struct {
		now time.Time
	}
	tests := []struct {
		name string
		c    *StandardDateCounter
		args args
		want *StandardDateCounter
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

			tt.c.rotate(tt.args.now)
		})
	}
}
