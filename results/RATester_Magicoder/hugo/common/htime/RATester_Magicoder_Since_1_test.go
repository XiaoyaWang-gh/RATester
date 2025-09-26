package htime

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSince_1(t *testing.T) {
	type args struct {
		time time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Duration
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

			if got := Since(tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Since() = %v, want %v", got, tt.want)
			}
		})
	}
}
