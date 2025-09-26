package web

import (
	"fmt"
	"testing"
	"time"
)

func TestDate_1(t *testing.T) {
	type args struct {
		t      time.Time
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				t:      time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
				format: "2006-01-02",
			},
			want: "2022-01-01",
		},
		{
			name: "Test case 2",
			args: args{
				t:      time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
				format: "2006-01-02 15:04:05",
			},
			want: "2022-01-01 12:00:00",
		},
		{
			name: "Test case 3",
			args: args{
				t:      time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
				format: "2006-01-02 15:04:05.000",
			},
			want: "2022-01-01 12:00:00.000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Date(tt.args.t, tt.args.format); got != tt.want {
				t.Errorf("Date() = %v, want %v", got, tt.want)
			}
		})
	}
}
