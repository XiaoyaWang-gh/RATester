package metric

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewStandardDateCounter_1(t *testing.T) {
	type args struct {
		reserveDays int64
	}
	tests := []struct {
		name string
		args args
		want *StandardDateCounter
	}{
		{
			name: "test1",
			args: args{
				reserveDays: 10,
			},
			want: &StandardDateCounter{
				reserveDays:    10,
				counts:         make([]int64, 10),
				lastUpdateDate: time.Now(),
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

			if got := newStandardDateCounter(tt.args.reserveDays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStandardDateCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}
