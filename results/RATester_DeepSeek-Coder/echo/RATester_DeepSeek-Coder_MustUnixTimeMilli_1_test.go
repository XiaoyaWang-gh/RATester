package echo

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestMustUnixTimeMilli_1(t *testing.T) {
	type args struct {
		sourceParam string
		dest        *time.Time
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
			if got := b.MustUnixTimeMilli(tt.args.sourceParam, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustUnixTimeMilli() = %v, want %v", got, tt.want)
			}
		})
	}
}
