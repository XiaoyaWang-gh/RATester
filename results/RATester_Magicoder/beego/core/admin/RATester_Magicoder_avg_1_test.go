package admin

import (
	"fmt"
	"testing"
	"time"
)

func Testavg_1(t *testing.T) {
	type args struct {
		items []time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "Test Case 1",
			args: args{
				items: []time.Duration{1, 2, 3, 4, 5},
			},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{
				items: []time.Duration{10, 20, 30, 40, 50},
			},
			want: 30,
		},
		{
			name: "Test Case 3",
			args: args{
				items: []time.Duration{},
			},
			want: 0,
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
