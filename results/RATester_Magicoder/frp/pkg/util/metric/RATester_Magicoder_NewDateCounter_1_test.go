package metric

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewDateCounter_1(t *testing.T) {
	type args struct {
		reserveDays int64
	}
	tests := []struct {
		name string
		args args
		want DateCounter
	}{
		{
			name: "TestNewDateCounter_0",
			args: args{
				reserveDays: 1,
			},
			want: newStandardDateCounter(1),
		},
		{
			name: "TestNewDateCounter_1",
			args: args{
				reserveDays: 0,
			},
			want: newStandardDateCounter(1),
		},
		{
			name: "TestNewDateCounter_2",
			args: args{
				reserveDays: -1,
			},
			want: newStandardDateCounter(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewDateCounter(tt.args.reserveDays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDateCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}
