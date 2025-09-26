package ratelimit

import (
	"fmt"
	"reflect"
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
		want bucketOption
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

			if got := withRate(tt.args.rate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("withRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
