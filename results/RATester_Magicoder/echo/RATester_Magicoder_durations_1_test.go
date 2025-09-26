package echo

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDurations_1(t *testing.T) {
	type args struct {
		sourceParam string
		values      []string
		dest        *[]time.Duration
	}
	tests := []struct {
		name string
		args args
		want *ValueBinder
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

			b := &ValueBinder{}
			if got := b.durations(tt.args.sourceParam, tt.args.values, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("durations() = %v, want %v", got, tt.want)
			}
		})
	}
}
