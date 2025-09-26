package framework

import (
	"fmt"
	"testing"
	"time"
)

func TestNowStamp_1(t *testing.T) {
	type args struct {
		time time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test nowStamp with current time",
			args: args{
				time: time.Now(),
			},
			want: time.Now().Format(time.StampMilli),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := nowStamp(); got != tt.want {
				t.Errorf("nowStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
