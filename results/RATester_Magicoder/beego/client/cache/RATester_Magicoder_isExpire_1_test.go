package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestisExpire_1(t *testing.T) {
	type args struct {
		mi *MemoryItem
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				mi: &MemoryItem{
					val:         "test",
					createdTime: time.Now(),
					lifespan:    1 * time.Second,
				},
			},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{
				mi: &MemoryItem{
					val:         "test",
					createdTime: time.Now().Add(-2 * time.Second),
					lifespan:    1 * time.Second,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.mi.isExpire(); got != tt.want {
				t.Errorf("isExpire() = %v, want %v", got, tt.want)
			}
		})
	}
}
