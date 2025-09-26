package echo

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDuration_1(t *testing.T) {
	type args struct {
		sourceParam    string
		dest           *time.Duration
		valueMustExist bool
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
			if got := b.duration(tt.args.sourceParam, tt.args.dest, tt.args.valueMustExist); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("duration() = %v, want %v", got, tt.want)
			}
		})
	}
}
