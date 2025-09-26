package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestWithRate_1(t *testing.T) {
	type args struct {
		rate time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test withRate",
			args: args{
				rate: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			withRate(tt.args.rate)
		})
	}
}
