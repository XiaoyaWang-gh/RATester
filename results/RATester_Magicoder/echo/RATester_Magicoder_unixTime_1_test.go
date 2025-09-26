package echo

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestUnixTime_1(t *testing.T) {
	type args struct {
		sourceParam    string
		dest           *time.Time
		valueMustExist bool
		precision      time.Duration
	}
	tests := []struct {
		name string
		args args
		want *time.Time
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
			if got := b.unixTime(tt.args.sourceParam, tt.args.dest, tt.args.valueMustExist, tt.args.precision); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unixTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
