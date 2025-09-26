package alils

import (
	"fmt"
	"testing"
	"time"
)

func TestNowRFC1123_1(t *testing.T) {
	type args struct {
		time time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test nowRFC1123 with current time",
			args: args{
				time: time.Now(),
			},
			want: time.Now().In(gmtLoc).Format(time.RFC1123),
		},
		{
			name: "Test nowRFC1123 with specific time",
			args: args{
				time: time.Date(2022, 1, 1, 10, 0, 0, 0, gmtLoc),
			},
			want: "Mon, 01 Jan 2022 10:00:00 GMT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := nowRFC1123(); got != tt.want {
				t.Errorf("nowRFC1123() = %v, want %v", got, tt.want)
			}
		})
	}
}
