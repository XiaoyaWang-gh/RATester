package ratelimit

import (
	"fmt"
	"reflect"
	"testing"
)

func TestwithCapacity_1(t *testing.T) {
	type args struct {
		capacity uint
	}
	tests := []struct {
		name string
		args args
		want bucketOption
	}{
		{
			name: "Test with capacity 10",
			args: args{capacity: 10},
			want: withCapacity(10),
		},
		{
			name: "Test with capacity 20",
			args: args{capacity: 20},
			want: withCapacity(20),
		},
		{
			name: "Test with capacity 30",
			args: args{capacity: 30},
			want: withCapacity(30),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := withCapacity(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("withCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}
