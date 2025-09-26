package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestGetRate_1(t *testing.T) {
	type fields struct {
		remaining   uint
		capacity    uint
		lastCheckAt time.Time
		rate        time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "test getRate",
			fields: fields{
				remaining:   10,
				capacity:    10,
				lastCheckAt: time.Now(),
				rate:        10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b := &tokenBucket{
				remaining:   tt.fields.remaining,
				capacity:    tt.fields.capacity,
				lastCheckAt: tt.fields.lastCheckAt,
				rate:        tt.fields.rate,
			}
			if got := b.getRate(); got != tt.want {
				t.Errorf("getRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
