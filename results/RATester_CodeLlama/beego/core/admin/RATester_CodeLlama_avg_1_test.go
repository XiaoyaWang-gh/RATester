package admin

import (
	"fmt"
	"testing"
	"time"
)

func TestAvg_1(t *testing.T) {
	type args struct {
		items []time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "test1",
			args: args{
				items: []time.Duration{
					time.Second * 1,
					time.Second * 2,
					time.Second * 3,
				},
			},
			want: time.Second * 2,
		},
		{
			name: "test2",
			args: args{
				items: []time.Duration{
					time.Second * 1,
					time.Second * 2,
					time.Second * 3,
					time.Second * 4,
				},
			},
			want: time.Second * 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := avg(tt.args.items); got != tt.want {
				t.Errorf("avg() = %v, want %v", got, tt.want)
			}
		})
	}
}
